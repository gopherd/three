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
	mesh.Init()
	mesh.geometry = geometry
	mesh.material = material
	return mesh
}

// Bounds implements Object Bounds method
func (mesh *Mesh) Bounds() (min, max core.Vector3, ok bool) {
	min, max = mesh.geometry.Bounds()
	return min, max, min != max
}

// Render implements Object Render method
func (mesh *Mesh) Render(renderer renderer.Renderer, cameraTransform, transform core.Matrix4) {
	mesh.object3d.Render(renderer, cameraTransform, transform)

	var shader = mesh.material.Shader()
	if !mesh.object3d.program.created && !mesh.object3d.program.fail {
		if err := mesh.object3d.createProgram(renderer, shader); err != nil {
			panic(err)
		}
	}
	var needsUpdate = mesh.material.NeedsUpdate()
	if needsUpdate {
		mesh.material.SetNeedsUpdate(false)
		for name, uniform := range shader.Uniforms {
			renderer.SetUniform(mesh.object3d.program.Id, name, uniform)
		}
	}

	var attributes = mesh.geometry.Attributes()
	var index = mesh.geometry.Index()
	if mesh.geometry.NeedsUpdate() {
		mesh.geometry.SetNeedsUpdate(false)
		// TODO: update attributes
	}
	_, _ = attributes, index
}
