package application

import (
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/driver/window"
	"github.com/gopherd/threego/three/object"
)

type Application struct {
	window   window.Window
	renderer renderer.Renderer
	scenes   []object.Scene
	camera   object.Camera
}

func New() *Application {
	return &Application{}
}

// Init implements three.Application Init method
func (app *Application) Init(window window.Window, renderder renderer.Renderer) error {
	app.window = window
	app.renderer = renderer
	return nil
}

// Shutdown implements three.Application Shutdown method
func (app *Application) Shutdown() {
}

// Update implements three.Application Update method
func (app *Application) Update() {
	var scene = app.GetRunningScene()
	if scene == nil {
		return
	}
	object.Update(scene)
	if app.camera != nil {
		scene.Render(app.renderer, app.camera)
	}
}

func (app *Application) GetCamera() object.Camera {
	return app.camera
}

func (app *Application) SetCamera(camera object.Camera) {
	app.camera = camera
}

func (app *Application) GetRunningScene() object.Scene {
	if n := len(app.scenes); n > 0 {
		return app.scenes[n-1]
	}
	return nil
}

func (app *Application) RunScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
		app.scenes = app.scenes[:0]
	}
	app.PushScene(scene)
}

func (app *Application) PushScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
	}
	app.scenes = append(app.scenes, scene)
	scene.OnEnter()
}

func (app *Application) PopScene(scene object.Scene) {
	if n := len(app.scenes); n > 0 {
		app.scenes[n-1].OnExit()
		app.scenes = app.scenes[:n-1]
		if n > 1 {
			app.scenes[n-2].OnEnter()
		}
	}
}
