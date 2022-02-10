package window

import "github.com/gopherd/threego/three/driver/renderer"

type Window interface {
	Init(renderer renderer.Renderer, title string, width, height int) error
	Terminate()
	Update()
	ShouldClose() bool
}
