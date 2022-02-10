package renderer

type Uniform interface{}

type Renderer interface {
	Init(width, height int) error
	Viewport(x, y, w, h int32)
	ClearColor(r, g, b, a float32)
	CreateProgram(vshader, fshader string) (uint32, error)
	LinkProgram(program uint32) error
	SetUniform(program uint32, name string, uniform Uniform)
}
