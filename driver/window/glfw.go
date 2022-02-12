package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/gopherd/three/driver/renderer"
)

type glfwindow struct {
	window *glfw.Window
}

func GLFWindow() Window {
	return &glfwindow{}
}

func (w *glfwindow) Init(renderer renderer.Renderer, title string, width, height int) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()
	w.window = window
	w.window.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
		renderer.Viewport(0, 0, int32(width), int32(height))
	})
	return nil
}

func (w *glfwindow) Terminate() {
	glfw.Terminate()
}

func (w *glfwindow) Update() {
	w.window.SwapBuffers()
	glfw.PollEvents()
}

func (w *glfwindow) ShouldClose() bool {
	return w.window.ShouldClose()
}
