package boot

import (
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"

	"github.com/gopherd/doge/query"
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
	options.Title = query.Or(options.Title, "Title")
	options.Width = query.Or(options.Width, 800)
	options.Height = query.Or(options.Height, 600)
	options.Window = query.OrNew(options.Window, window.GLFWindow)
	options.Renderer = query.OrNew(options.Renderer, renderer.OpenGLRenderer)
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
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for sig := range sigChan {
			if sig == os.Interrupt {
				break
			}
		}
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
