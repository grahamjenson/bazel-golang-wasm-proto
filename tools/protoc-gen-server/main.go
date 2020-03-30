package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"log"
	"text/template"
)

// main setup and render the generated code
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

// protoModule is the struct used to generate the code
type protoModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func (m *protoModule) Name() string { return "server" }

func (m *protoModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
	tpl := template.New("server")
	m.tpl = template.Must(tpl.Parse(serviceTpl))
}

func (m *protoModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		if len(f.Services()) == 0 {
			continue
		}
		name := m.ctx.OutputPath(f).SetExt(".server.go")
		m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
	}
	return m.Artifacts()
}

// the code template
const serviceTpl = `package {{ .Package.ProtoName }}

import (
	"context"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

{{ range .Services }}
func Register{{ .Name }}HTTPMux(mux *http.ServeMux, srv {{ .Name }}Server) {
	{{ range .Methods}}
		{{ $method := printf "/%s.%s/%s" .Service.Package.ProtoName .Service.Name .Name }}
	mux.HandleFunc("{{ $method }}", func(w http.ResponseWriter, r *http.Request) {
		in := new({{ .Input.Name }})
		inJSON, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.Unmarshal(inJSON, in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		ret, err := srv.{{ .Name }}(context.Background(), in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return 
		}
		retJSON, err := json.Marshal(ret)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return 
		}
		w.Write(retJSON)
	})
	{{ end }}
}
{{ end }}
`
