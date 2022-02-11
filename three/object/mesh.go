package object

import (
	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/geometry"
	"github.com/gopherd/threego/three/math"
)

var _ Object = (*Mesh)(nil)

type Mesh struct {
	object3d
	geometry geometry.Geometry
}

// TODO: Bounds implements Object Bounds method
func (mesh *Mesh) Bounds() (min, max math.Vector3, ok bool) {
	return
}

// TODO: Render implements Object Render method
func (mesh *Mesh) Render(renderer renderer.Renderer, cameraTransform, transform math.Mat4x4) {
}
