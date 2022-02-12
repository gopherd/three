package boot

import (
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"

	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/driver/window"
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
	Title  string
	Width  int
	Height int

	Window   window.Window
	Renderer renderer.Renderer

	Start func()
}

func (options *Options) init() {
	if options.Title == "" {
		options.Title = "Title"
	}
	if options.Width <= 0 {
		options.Width = 800
	}
	if options.Height <= 0 {
		options.Height = 600
	}
	if options.Window == nil {
		options.Window = window.GLFWindow()
	}
	if options.Renderer == nil {
		options.Renderer = renderer.OpenGLRenderer()
	}
}

func Run(app Application, options Options) {
	options.init()
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

	if options.Start != nil {
		options.Start()
	}

	for !options.Window.ShouldClose() && atomic.LoadInt32(&quit) == 0 {
		app.Update()
		options.Window.Update()
	}
}
