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

func NewMesh(geometry geometry.Geometry, material material.Material) *Mesh {
	var mesh = new(Mesh)
	mesh.geometry = geometry
	mesh.material = material
	mesh.Init()
	return mesh
}

func (mesh *Mesh) String() string {
	return "Mesh.String:TODO"
}

// Bounds implements Object Bounds method
func (mesh *Mesh) Bounds() geometry.Box3 {
	return mesh.geometry.Bounds()
}

// Render implements Object Render method
func (mesh *Mesh) Render(renderer renderer.Renderer, proj, view, transform core.Matrix4) {
	mesh.object3d.Render(renderer, proj, view, transform)
	mesh.object3d.renderGeometry(renderer, mesh.geometry, mesh.material)
}
