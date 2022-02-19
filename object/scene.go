package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
)

// Scene represents a scene graph should be rendered
type Scene interface {
	node

	// Add adds object to the scene
	Add(object Object)
	// Render renders the scene by camera to renderer
	Render(renderer renderer.Renderer, camera Camera)

	// OnEnter callback called on enter the scene
	OnEnter()
	// OnExit callback called on exit the scene
	OnExit()
}

// Update updates scene
func Update(scene Scene) {
	recursivelyUpdateNode(scene)
}

var _ Scene = (*BasicScene)(nil)

// BasicScene implements a basic Scene
type BasicScene struct {
	node3d
	background core.Vector4
}

func (scene *BasicScene) ToString() string {
	return "Scene"
}

// SetBackground sets the scene background color
func (scene *BasicScene) SetBackground(color core.Vector4) {
	scene.background = color
}

// Add implements Scene Add method
func (scene *BasicScene) Add(object Object) {
	scene.addChild(object)
}

// Render implements Scene Render method
func (scene *BasicScene) Render(renderer renderer.Renderer, camera Camera) {
	var proj = camera.Projection()
	var view = camera.TransformWorld()
	var background = scene.background
	renderer.ClearColor(background.X(), background.Y(), background.Z(), background.W())
	for _, child := range scene.children {
		if !child.Visible() {
			continue
		}
		recursivelyRenderObject(renderer, camera, proj, view, child, child.Transform())
	}
}

// OnEnter implements Scene OnEnter method
func (scene *BasicScene) OnEnter() {}

// OnExit implements Scene OnExit method
func (scene *BasicScene) OnExit() {}
