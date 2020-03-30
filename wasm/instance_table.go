package main

import (
	"fmt"
	"github.com/grahamjenson/bazel-golang-wasm-protoc/protos/api"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type InstanceTable struct {
	app.Compo
	manager   *Manager
	instances []*api.Instance
}

func (p *InstanceTable) SetManager(manager *Manager) {
	p.manager = manager
}

func (p *InstanceTable) Render() app.UI {

	nodes := []app.Node{}
	for _, i := range p.instances {
		nodes = append(nodes, app.Tr().Body(
			app.Td().Body(app.Text(i.Name)),
			app.Td().Body(app.Text(i.InstanceType)),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.Ecu))),
			app.Td().Body(app.Text(fmt.Sprintf("%v", i.Memory))),
			app.Td().Body(app.Text(i.Network)),
			app.Td().Body(app.Text(i.Price)),
		))
	}

	return app.Table().Class("table").Body(
		app.Tr().Body(
			app.Th().Scope("col").Body(app.Text("Name")),
			app.Th().Scope("col").Body(app.Text("Instance Type")),
			app.Th().Scope("col").Body(app.Text("ECU")),
			app.Th().Scope("col").Body(app.Text("Mem")),
			app.Th().Scope("col").Body(app.Text("Network")),
			app.Th().Scope("col").Body(app.Text("Price")),
		),
		app.TBody().Body(nodes...),
	)

}
