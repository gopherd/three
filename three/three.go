package three

import (
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"

	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/driver/window"
)

func init() {
	runtime.LockOSThread()
}

type Application interface {
	Init(window window.Window, renderer renderer.Renderer) error
	Update()
	Shutdown()
}

type Options struct {
	Title    string
	Width    int
	Height   int
	Window   window.Window
	Renderer renderer.Renderer
}

func DefaultOptions(title string) Options {
	return Options{
		Title:    title,
		Width:    800,
		Height:   600,
		Window:   window.GLFWindow(),
		Renderer: renderer.OpenGLRenderer(),
	}
}

func Run(app Application, options Options) {
	if err := options.Window.Init(
		options.Renderer,
		options.Title,
		options.Width,
		options.Height,
	); err != nil {
		panic(err)
	}
	defer options.Window.Terminate()

	if err := options.Renderer.Init(options.Width, options.Height); err != nil {
		panic(err)
	}

	if err := app.Init(options.Window, options.Renderer); err != nil {
		panic(err)
	}
	defer app.Shutdown()

	var quit int32
	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		atomic.StoreInt32(&quit, 1)
	}()

	for !options.Window.ShouldClose() && atomic.LoadInt32(&quit) == 0 {
		app.Update()
		options.Window.Update()
	}
}
