package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
)

type Scene interface {
	node

	Add(object Object) bool
	Render(renderer renderer.Renderer, camera Camera)

	OnEnter()
	OnExit()
}

// Update updates scene
func Update(scene Scene) {
	recursivelyUpdateNode(scene)
}

var _ Scene = (*BasicScene)(nil)

type BasicScene struct {
	node3d
	background core.Vector4
}

func (scene *BasicScene) SetBackground(color core.Vector4) {
	scene.background = color
}

func (scene *BasicScene) Add(object Object) bool {
	return scene.addChild(object)
}

// Render implements Scene Render method
func (scene *BasicScene) Render(renderer renderer.Renderer, camera Camera) {
	var background = scene.background
	renderer.ClearColor(background.R(), background.G(), background.B(), background.A())
	var cameraTransform = TransformToWorld(camera)
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
