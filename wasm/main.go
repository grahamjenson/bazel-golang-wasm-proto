package main

import (
	"fmt"

	"github.com/grahamjenson/bazel-golang-wasm-proto/protos/api"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	manager := &Manager{
		searchBar:     &SearchBar{},
		instanceTable: &InstanceTable{},
	}

	manager.searchBar.SetManager(manager)
	manager.instanceTable.SetManager(manager)

	app.Route("/", manager)
	app.Run()
}

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo
	searchBar     *SearchBar
	instanceTable *InstanceTable
}

func (h *Manager) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.Nav().Class("navbar navbar-expand-lg navbar-light bg-light").Body(
				h.searchBar,
			),
		),
		app.Div().Class("container-fluid").Body(
			h.instanceTable,
		),
	)
}

func (h *Manager) Search(q string) []*api.Instance {
	instances, err := api.CallApiSearch(api.SearchRequest{
		Query: q,
	})

	if err != nil {
		fmt.Println("Search Error:", err)
		return []*api.Instance{}
	}

	return instances.Instances
}

func (h *Manager) UpdateInstances(q string) {
	instances := h.Search(q)
	h.instanceTable.instances = instances
	h.instanceTable.Update()
}
