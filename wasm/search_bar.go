package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SearchBar struct {
	app.Compo
	// Must be a pointer, or else the widget construction will end up incorrect.
	Manager      *Manager
	searchString string
}

func (p *SearchBar) SetManager(manager *Manager) {
	p.Manager = manager
	if p.Manager == nil {
		panic("p.Manager == nil")
	}
}

func (p *SearchBar) Render() app.UI {
	input := app.Input().
		Class("form-control").
		Value(p.searchString).
		Placeholder("t2.small").
		AutoFocus(true).
		OnKeyUp(p.OnInputChange)

	return app.Div().Class("input-group").Body(
		app.
			Div().
			Class("input-group-prepend").
			Body(app.
				Span().
				Class("input-group-text").
				Body(app.Text("🔍"))),
		input,
	)
}

func (p *SearchBar) OnInputChange(ctx app.Context, e app.Event) {
	if p == nil {
		panic("Manager == nil, why?")
	}
	src := ctx.JSSrc()
	p.searchString = src.Get("value").String()
	p.Update()
	if p.Manager == nil {
		panic("p.Manager == nil")
	}
	p.Manager.UpdateInstances(p.searchString)
}
