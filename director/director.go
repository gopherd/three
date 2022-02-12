package director

import (
	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/driver/window"
	"github.com/gopherd/three/object"
)

type Director struct {
	window   window.Window
	renderer renderer.Renderer
	scenes   []object.Scene
	camera   object.Camera
}

var app = new(Director)

func Get() *Director {
	return app
}

// Init implements boot.Application Init method
func (app *Director) Init(window window.Window, renderer renderer.Renderer) error {
	app.window = window
	app.renderer = renderer
	return nil
}

// Shutdown implements boot.Application Shutdown method
func (app *Director) Shutdown() {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
		app.scenes = app.scenes[:0]
	}
}

// Update implements boot.Application Update method
func (app *Director) Update() {
	var scene = app.GetRunningScene()
	if scene == nil {
		return
	}
	object.Update(scene)
	if app.camera != nil {
		scene.Render(app.renderer, app.camera)
	}
}

func (app *Director) GetCamera() object.Camera {
	return app.camera
}

func (app *Director) SetCamera(camera object.Camera) {
	app.camera = camera
}

func (app *Director) GetRunningScene() object.Scene {
	if n := len(app.scenes); n > 0 {
		return app.scenes[n-1]
	}
	return nil
}

func (app *Director) RunScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
		app.scenes = app.scenes[:0]
	}
	app.PushScene(scene)
}

func (app *Director) PushScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
	}
	app.scenes = append(app.scenes, scene)
	scene.OnEnter()
}

func (app *Director) PopScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
		app.scenes = app.scenes[:n-1]
		if n > 1 {
			app.scenes[n-2].OnEnter()
		}
	}
}
