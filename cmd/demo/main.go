package main

import (
	"log"

	"github.com/gopherd/three/boot"
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/director"
	"github.com/gopherd/three/object"
)

func main() {
	app := director.Get()
	app.RunScene(NewScene())
	boot.Run(app, boot.DefaultOptions("Demo"))
}

type Scene struct {
	object.BasicScene
	camera object.Camera
}

func NewScene() *Scene {
	scene := new(Scene)
	scene.SetBackground(core.Vec4(0.2, 0.3, 0.3, 1.0))
	scene.camera = object.NewPerspectiveCamera()
	scene.Add(scene.camera)
	director.Get().SetCamera(scene.camera)
	return scene
}

func (scene *Scene) OnEnter() {
	log.Println("scene.OnEnter")
}

func (scene *Scene) OnExit() {
	log.Println("scene.OnExit")
}
