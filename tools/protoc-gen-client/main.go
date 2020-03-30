package main

import (
	"log"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type protoModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func (m *protoModule) Name() string { return "client" }

func (m *protoModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("client")
	m.tpl = template.Must(tpl.Parse(serviceTpl))
}

func (m *protoModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		if len(f.Services()) == 0 {
			continue
		}

		name := m.ctx.OutputPath(f).SetExt(".client.go")
		m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
	}

	return m.Artifacts()
}

func main() {
	log.SetFlags(0)
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		&protoModule{ModuleBase: &pgs.ModuleBase{}},
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}

const serviceTpl = `package {{ .Package.ProtoName }}

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
)

{{ range .Services }}
{{ range .Methods}}

func Call{{ .Service.Name }}{{ .Name }}(input {{ .Input.Name }}) (*{{ .Output.Name }}, error) {
	{{ $method := printf "/%s.%s/%s" .Service.Package.ProtoName .Service.Name .Name }}

	str, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "{{ $method }}", strings.NewReader(string(str)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := {{ .Output.Name }}{}
	err = json.Unmarshal(body, &instances)
	if err != nil {
		return nil, err
	}

	return &instances, nil
}
{{ end }}
{{ end }}
`
