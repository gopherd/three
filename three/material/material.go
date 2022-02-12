package material

import "github.com/gopherd/threego/three/driver/renderer/shader"

type FaceSide int

const (
	FrontSide FaceSide = iota
	BackSide
	DoubleSide
)

type Options struct {
	Side         FaceSide
	Transparent  bool
	Opacity      float32
	VertexColors float32
}

type Material interface {
	Options() Options
	Shader() shader.Shader
}

type material struct {
	options Options
	shader  shader.Shader
}

func (m *material) Opitions() Options {
	return m.options
}

func (m *material) Shader() shader.Shader {
	return m.shader
}
