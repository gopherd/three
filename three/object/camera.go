package object

import (
	"github.com/gopherd/doge/math/tensor"

	"github.com/gopherd/threego/three/driver/renderer"
)

// Camera represents camera object
type Camera interface {
	Object
	ContainsBox(transform tensor.Mat4x4, min, max tensor.Vector3) bool
}

// PerspectiveCamera implemets perspective camera
type PerspectiveCamera struct {
	object3d
}

var _ Camera = (*PerspectiveCamera)(nil)

// NewPerspectiveCamera creates a PerspectiveCamera
func NewPerspectiveCamera() *PerspectiveCamera {
	camera := new(PerspectiveCamera)
	camera.Init()
	return camera
}

// TODO(delay) Bounds implements Object Bounds method
func (camera *PerspectiveCamera) Bounds() (min, max tensor.Vector3, ok bool) {
	return
}

// TODO(delay) Render implements Object Render method
func (camera *PerspectiveCamera) Render(renderer renderer.Renderer, cameraTransform, transform tensor.Mat4x4) {
}

// TODO(delay) ContainsBox implements Camera ContainsBox method
func (camera *PerspectiveCamera) ContainsBox(transform tensor.Mat4x4, min, max tensor.Vector3) bool {
	return true
}
