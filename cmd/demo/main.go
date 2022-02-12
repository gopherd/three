package main

import (
	"log"

	"github.com/gopherd/doge/math/tensor"

	"github.com/gopherd/threego/three"
	"github.com/gopherd/threego/three/application"
	"github.com/gopherd/threego/three/object"
)

func main() {
	app := application.Get()
	app.RunScene(NewScene())
	three.Run(app, three.DefaultOptions("Demo"))
}

type Scene struct {
	object.BasicScene
	camera object.Camera
}

func NewScene() *Scene {
	scene := new(Scene)
	scene.camera = object.NewPerspectiveCamera()
	scene.Add(scene.camera)
	application.Get().SetCamera(scene.camera)
	scene.SetBackground(tensor.Vec4(0.2, 0.3, 0.3, 1.0))
	return scene
}

func (scene *Scene) OnEnter() {
	log.Println("scene.OnEnter")
}

func (scene *Scene) OnExit() {
	log.Println("scene.OnExit")
}
