package object

import (
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/math"
)

type Scene interface {
	node
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
	background math.Vector4
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

func (scene *BasicScene) SetBackground(color math.Vector4) {
	scene.background = color
}

// OnEnter implements Scene OnEnter method
func (scene *BasicScene) OnEnter() {}

// OnExit implements Scene OnExit method
func (scene *BasicScene) OnExit() {}
