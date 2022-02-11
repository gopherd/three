package geometry

import (
	"github.com/gopherd/threego/three/math"
)

type Range struct {
	Start int
	End   int
}

type Group struct {
	Range         Range
	MaterialIndex int
}

type Geometry interface {
	Index() Attribute
	SetIndex(Attribute)
	GetAttribute(name string) Attribute
	SetAttribute(name string, attribute Attribute)
	Bounds() (min, max math.Vector3)
	DrawRange() Range
	SetDrawRange(Range)
	AddGroup(Group)
	Groups() []Group
}

var _ Geometry = (*BufferGeometry)(nil)

type BufferGeometry struct {
	indices    Attribute
	attributes map[string]Attribute
	bounds     struct{ min, max math.Vector3 }
	drawRange  Range
	groups     []Group
}

func (geo *BufferGeometry) Index() Attribute {
	return geo.indices
}

func (geo *BufferGeometry) SetIndex(indices Attribute) {
	geo.indices = indices
}

func (geo *BufferGeometry) GetAttribute(name string) Attribute {
	return geo.attributes[name]
}

func (geo *BufferGeometry) SetAttribute(name string, attribute Attribute) {
	geo.attributes[name] = attribute
}

func (geo *BufferGeometry) Bounds() (min, max math.Vector3) {
	return geo.bounds.min, geo.bounds.max
}

func (geo *BufferGeometry) DrawRange() Range {
	return geo.drawRange
}

func (geo *BufferGeometry) SetDrawRange(drawRange Range) {
	geo.drawRange = drawRange
}

func (geo *BufferGeometry) AddGroup(group Group) {
	geo.groups = append(geo.groups, group)
}

func (geo *BufferGeometry) Groups() []Group {
	return geo.groups
}
