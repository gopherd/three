package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/geometry"
	"github.com/gopherd/three/material"
)

var _ Object = (*Mesh)(nil)

type Mesh struct {
	object3d
	geometry geometry.Geometry
	material material.Material
}

// TODO: Bounds implements Object Bounds method
func (mesh *Mesh) Bounds() (min, max core.Vector3, ok bool) {
	return
}

// TODO: Render implements Object Render method
func (mesh *Mesh) Render(renderer renderer.Renderer, cameraTransform, transform core.Mat4x4) {
}
