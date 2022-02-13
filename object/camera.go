package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
)

// Camera represents a camera object
type Camera interface {
	Object
	ContainsBox(transform core.Mat4x4, min, max core.Vector3) bool
}

type cameraImpl struct {
	object3d
}

// TODO(delay) Bounds implements Object Bounds method
func (camera *cameraImpl) Bounds() (min, max core.Vector3, ok bool) {
	return
}

// TODO(delay) Render implements Object Render method
func (camera *cameraImpl) Render(renderer renderer.Renderer, cameraTransform, transform core.Mat4x4) {
}

// PerspectiveCamera represents a perspective camera
type PerspectiveCamera interface {
	Camera
	isPerspectiveCamera()
}

// perspectiveCamera implemets perspective camera
type perspectiveCameraImpl struct {
	cameraImpl
}

var _ PerspectiveCamera = (*perspectiveCameraImpl)(nil)

// NewPerspectiveCamera creates a PerspectiveCamera
func NewPerspectiveCamera() PerspectiveCamera {
	camera := new(perspectiveCameraImpl)
	camera.Init()
	return camera
}

// isPerspectiveCamera implements PerspectiveCamera isPerspectiveCamera method
func (camera *perspectiveCameraImpl) isPerspectiveCamera() {}

// TODO(delay) ContainsBox implements Camera ContainsBox method
func (camera *perspectiveCameraImpl) ContainsBox(transform core.Mat4x4, min, max core.Vector3) bool {
	return true
}
