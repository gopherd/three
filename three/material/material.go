package material

import (
	"github.com/gopherd/threego/three/driver/renderer"
)

type FaceSide int

const (
	FrontSide FaceSide = iota
	BackSide
	DoubleSide
)

type Shader interface {
	Uniforms() map[string]renderer.Uniform
	VertexShader() string
	FragmentShader() string
}

type Material interface {
	Side() FaceSide
	Transparent() bool
	Opacity() float32
	VertexColors() bool
	Shader() Shader
}
