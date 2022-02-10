package object

import (
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/math"
)

type Scene struct {
	node
}

func (scene *Scene) Render(renderer renderer.Renderer, camera Camera) {
	cameraTransform := camera.Transform()
	for _, child := range scene.children {
		recursivelyRenderObject(renderer, cameraTransform, child.Transform(), child)
	}
}

func recursivelyRenderObject(renderer renderer.Renderer, camera, transform math.Mat4x4, obj Object) {
	obj.Render(renderer, camera, transform)
	for i, n := 0, obj.NumChild(); i < n; i++ {
		child := obj.GetChildByIndex(i)
		recursivelyRenderObject(renderer, camera, transform.Mul(child.Transform()), child)
	}
}
