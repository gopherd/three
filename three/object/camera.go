package object

import (
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/math"
)

// Camera represents camera object
type Camera interface {
	Object
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

// Render implements Object Render method
func (camera *PerspectiveCamera) Render(renderer renderer.Renderer, cameraTransform, transform math.Mat4x4) {
}
