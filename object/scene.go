package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
)

// A Scene represents a root node should be rendered
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
	var background = scene.background
	renderer.ClearColor(background.R(), background.G(), background.B(), background.A())
	var cameraTransform = camera.WorldTransform()
	for _, child := range scene.children {
		if !child.Visible() {
			continue
		}
		recursivelyRenderObject(renderer, camera, cameraTransform, child, child.Transform())
	}
}

// OnEnter implements Scene OnEnter method
func (scene *BasicScene) OnEnter() {}

// OnExit implements Scene OnExit method
func (scene *BasicScene) OnExit() {}
