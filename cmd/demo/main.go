package main

import (
	"log"

	"github.com/gopherd/threego/three"
	"github.com/gopherd/threego/three/application"
	"github.com/gopherd/threego/three/math"
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
	object.Add(scene, scene.camera)
	application.Get().SetCamera(scene.camera)
	scene.SetBackground(math.Vec4(0.2, 0.3, 0.3, 1.0))
	return scene
}

func (scene *Scene) OnEnter() {
	log.Println("scene.OnEnter")
}

func (scene *Scene) OnExit() {
	log.Println("scene.OnExit")
}
