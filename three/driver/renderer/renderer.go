package renderer

type Renderer interface {
	Init(width, height int) error
	Viewport(x, y, w, h int32)
	ClearColor(r, g, b, a float32)
}
