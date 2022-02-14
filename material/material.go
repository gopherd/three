package material

import "github.com/gopherd/three/driver/renderer/shader"

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
	NeedsUpdate() bool
	SetNeedsUpdate(bool)
}

type basicMaterial struct {
	notNeedsUpdate bool
}

func (m *basicMaterial) NeedsUpdate() bool {
	return !m.notNeedsUpdate
}

func (m *basicMaterial) SetNeedsUpdate(needsUpdate bool) {
	m.notNeedsUpdate = !needsUpdate
}
