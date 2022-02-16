package main

import (
	"log"

	"github.com/gopherd/three/boot"
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/director"
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
}

func NewScene() *Scene {
	const aspect = 1
	object.AddedEventType = nil
	scene := new(Scene)
	scene.SetBackground(core.Vec4(0.2, 0.3, 0.3, 1.0))
	scene.camera = object.NewPerspectiveCamera(30, aspect, 0.1, 1000)
	scene.Add(scene.camera)
	director.SetCamera(scene.camera)
	return scene
}

func (scene *Scene) OnEnter() {
	log.Println("scene.OnEnter")
}

func (scene *Scene) OnExit() {
	log.Println("scene.OnExit")
}
