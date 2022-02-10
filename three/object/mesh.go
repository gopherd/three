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

func (mesh *Mesh) Render(renderer renderer.Renderer, camera, transform math.Mat4x4) {
}
