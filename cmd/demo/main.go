package main

import (
	"log"

	"github.com/gopherd/three/boot"
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/director"
	"github.com/gopherd/three/geometry"
	"github.com/gopherd/three/material"
	"github.com/gopherd/three/object"
)

func main() {
	boot.Run(director.Application, boot.Options{
		Title: "Demo",
		Start: func() {
			director.RunScene(NewScene())
		},
	})
}

type Scene struct {
	object.BasicScene
	camera object.Camera
	mesh   *object.Mesh
}

func NewScene() *Scene {
	scene := new(Scene)
	scene.SetBackground(core.Vec4(0.2, 0.3, 0.3, 1.0))

	// 设置相机
	const aspect = 1
	scene.camera = object.NewPerspectiveCamera(30, aspect, 0.1, 1000)
	scene.Add(scene.camera)
	director.SetCamera(scene.camera)

	// 添加 mesh
	scene.mesh = scene.createMesh()
	scene.Add(scene.mesh)

	print(object.Stringify(scene))

	return scene
}

func (scene *Scene) createMesh() *object.Mesh {
	var g = geometry.NewBufferGeometry()
	var positions = geometry.NewFloat32Attribute(3, 3)
	const z = 1
	positions.SetXYZ(0, 0, 0, 1)
	positions.SetXYZ(1, 1, 0, 1)
	positions.SetXYZ(2, 0, 1, 1)
	g.SetAttribute(geometry.AttributePosition, positions)
	var colors = geometry.NewFloat32Attribute(3, 3)
	colors.SetXYZ(0, 0, 0, 1)
	colors.SetXYZ(1, 1, 0, 0)
	colors.SetXYZ(2, 0, 1, 0)
	g.SetAttribute(geometry.AttributeColor, colors)
	g.ComputeBounds()
	var m = material.NewMeshBasicMaterial(material.MeshBasicMaterialParameters{
		Options: material.Options{
			VertexColors: true,
		},
	})
	return object.NewMesh(g, m)
}

func (scene *Scene) OnEnter() {
	log.Println("scene.OnEnter")
}

func (scene *Scene) OnExit() {
	log.Println("scene.OnExit")
}
