package main

import (
	"github.com/gopherd/threego/three"
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/driver/window"
)

func main() {
	three.Run(newApplication(), three.DefaultOptions("Demo"))
}

type application struct {
	window   window.Window
	renderer renderer.Renderer
}

func newApplication() *application {
	return new(application)
}

func (app *application) Init(window window.Window, renderer renderer.Renderer) error {
	app.window = window
	app.renderer = renderer
	return nil
}

func (app *application) Shutdown() {
}

func (app *application) Update() {
	app.renderer.ClearColor(0.2, 0.3, 0.3, 1.0)
}
