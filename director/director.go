package director

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gopherd/three/boot"
	"github.com/gopherd/three/core/event"
	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/driver/window"
	"github.com/gopherd/three/object"
)

var director struct {
	dispatcher event.Dispatcher
	window     window.Window
	renderer   renderer.Renderer

	scenes []object.Scene
	camera object.Camera

	updatedAt time.Time
	deltaTime time.Duration
}

var Application boot.Application = application{}

type application struct{}

// Init implements boot.Application Init method
func (application) Init(window window.Window, renderer renderer.Renderer) error {
	director.window = window
	director.renderer = renderer
	director.updatedAt = time.Now()
	return nil
}

// Shutdown implements boot.Application Shutdown method
func (application) Shutdown() {
	if n := len(director.scenes); n > 0 {
		director.scenes[n-1].OnExit()
		director.scenes = director.scenes[:0]
	}
}

// Update implements boot.Application Update method
func (application) Update() {
	defer func() {
		if e := recover(); e != nil {
			println(fmt.Sprintf("Error: %v\nStack:\n%v", e, string(debug.Stack())))
		}
	}()
	var now = time.Now()
	director.deltaTime = now.Sub(director.updatedAt)
	director.updatedAt = now

	var scene = GetRunningScene()
	if scene == nil {
		return
	}
	object.Update(scene)
	if director.camera != nil {
		scene.Render(director.renderer, director.camera)
	}
}

// DeltaTime returns delta time duration from last update
func DeltaTime() time.Duration {
	return director.deltaTime
}

// GetCamera returns current camera
func GetCamera() object.Camera {
	return director.camera
}

// SetCamera sets current camera
func SetCamera(camera object.Camera) {
	director.camera = camera
}

// GetRunningScene return the running scene
func GetRunningScene() object.Scene {
	if n := len(director.scenes); n > 0 {
		return director.scenes[n-1]
	}
	return nil
}

// RunScene runs the scene
func RunScene(scene object.Scene) {
	if n := len(director.scenes); n > 0 {
		director.scenes[n-1].OnExit()
		director.scenes = director.scenes[:0]
	}
	PushScene(scene)
}

// PushScene push and runs the scene
func PushScene(scene object.Scene) {
	if n := len(director.scenes); n > 0 {
		director.scenes[n-1].OnExit()
	}
	director.scenes = append(director.scenes, scene)
	scene.OnEnter()
}

// PopScene pops current running scene and runs previous scene
func PopScene(scene object.Scene) {
	if n := len(director.scenes); n > 0 {
		director.scenes[n-1].OnExit()
		director.scenes = director.scenes[:n-1]
		if n > 1 {
			director.scenes[n-2].OnEnter()
		}
	}
}

// Dispatcher returns the event dispatcher
func Dispatcher() *event.Dispatcher {
	return &director.dispatcher
}
