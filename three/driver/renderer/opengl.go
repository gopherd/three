package renderer

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type OpenGLRenderer struct {
}

func NewOpenGLRenderer() *OpenGLRenderer {
	return &OpenGLRenderer{}
}

func (OpenGLRenderer) Init(width, height int) error {
	if err := gl.Init(); err != nil {
		return err
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	return nil

}

func (OpenGLRenderer) Viewport(x, y, w, h int32) {
	gl.Viewport(x, y, w, h)
}

func (OpenGLRenderer) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
