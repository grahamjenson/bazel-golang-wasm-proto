package main

import (
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type SearchBar struct {
	app.Compo
	manager      *Manager
	searchString string
}

func (p *SearchBar) SetManager(manager *Manager) {
	p.manager = manager
}

func (p *SearchBar) Render() app.UI {
	input := app.Input().
		Class("form-control").
		Value(p.searchString).
		Placeholder("t2.small").
		AutoFocus(true).
		OnKeyup(p.OnInputChange)

	return app.Div().Class("input-group").Body(
		app.Div().Class("input-group-prepend").Body(app.Span().Class("input-group-text").Body(app.Text("🔍"))),
		input,
	)
}

func (p *SearchBar) OnInputChange(src app.Value, e app.Event) {
	p.searchString = src.Get("value").String()
	p.Update()
	p.manager.UpdateInstances(p.searchString)
}
