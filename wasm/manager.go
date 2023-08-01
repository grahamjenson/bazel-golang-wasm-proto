package main

import (
	"fmt"

	"github.com/grahamjenson/bazel-golang-wasm-proto/protos/api"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Manager is the main controller of this application, also the root Body
type Manager struct {
	app.Compo

	// These must be initialized in app.RenderFunc, or OnMount() for
	// Manager, and must be pointers.
	// See: https://github.com/maxence-charriere/go-app/issues/853.
	// This is not obvious.
	SearchBar     *SearchBar
	InstanceTable *InstanceTable
}

func (h *Manager) Render() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.
				Nav().
				Class("navbar navbar-expand-lg navbar-light bg-light").
				Body(h.SearchBar),
		),
		app.Div().Class("container-fluid").Body(h.InstanceTable),
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
	h.InstanceTable.instances = instances
	h.InstanceTable.Update()
}
