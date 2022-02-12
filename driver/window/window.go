package window

import "github.com/gopherd/three/driver/renderer"

type Window interface {
	Init(renderer renderer.Renderer, title string, width, height int) error
	Terminate()
	Update()
	ShouldClose() bool
}
