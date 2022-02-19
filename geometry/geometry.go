package geometry

import (
	"github.com/gopherd/doge/operator"
	"github.com/gopherd/three/core"
)

type Range struct {
	Start int
	End   int
}

type Group struct {
	Range         Range
	MaterialIndex int
}

type DrawPolicy int

const (
	StaticDraw DrawPolicy = iota
	DynamicDraw
	StreamDraw
)

type Geometry interface {
	Index() *Uint32Attribute
	Attributes() map[string]Attribute
	Bounds() Box3
	Groups() []Group
	DrawRange() Range
	DrawPolicy() DrawPolicy
	NeedsUpdate() bool
	SetNeedsUpdate(bool)
}

var _ Geometry = (*BufferGeometry)(nil)

type BufferGeometry struct {
	indices        *Uint32Attribute
	attributes     map[string]Attribute
	bounds         Box3
	groups         []Group
	drawRange      Range
	drawPolicy     DrawPolicy
	notNeedsUpdate bool
}

func NewBufferGeometry() *BufferGeometry {
	return &BufferGeometry{
		attributes: make(map[string]Attribute),
	}
}

func (geo *BufferGeometry) Index() *Uint32Attribute {
	return geo.indices
}

func (geo *BufferGeometry) SetIndex(indices *Uint32Attribute) {
	geo.indices = indices
}

func (geo *BufferGeometry) Attributes() map[string]Attribute {
	return geo.attributes
}

func (geo *BufferGeometry) GetAttribute(name string) Attribute {
	return geo.attributes[name]
}

func (geo *BufferGeometry) SetAttribute(name string, attribute Attribute) {
	geo.attributes[name] = attribute
}

func (geo *BufferGeometry) Bounds() Box3 {
	return geo.bounds
}

func (geo *BufferGeometry) ComputeBounds() bool {
	var positions, ok = geo.attributes[AttributePosition]
	if !ok {
		return false
	}
	var count = positions.Count()
	var stride = positions.Stride()
	if stride < 2 || stride > 3 {
		return false
	}
	var xmin, ymin, zmin core.Float
	var xmax, ymax, zmax core.Float
	for i := 0; i < count; i++ {
		var offset = i * stride
		var x, y, z core.Float
		x = positions.Float(offset)
		y = positions.Float(offset + 1)
		if stride == 3 {
			z = positions.Float(offset + 2)
		}
		xmin = operator.If(i == 0 || x < xmin, x, xmin)
		ymin = operator.If(i == 0 || y < ymin, y, ymin)
		zmin = operator.If(i == 0 || z < zmin, z, zmin)
		xmax = operator.If(i == 0 || x > xmax, x, xmax)
		ymax = operator.If(i == 0 || y > ymax, y, ymax)
		zmax = operator.If(i == 0 || z > zmax, z, zmax)
	}
	geo.bounds.Min = core.Vec3(xmin, ymin, zmin)
	geo.bounds.Max = core.Vec3(xmax, ymax, zmax)
	return true
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

func (geo *BufferGeometry) DrawPolicy() DrawPolicy {
	return geo.drawPolicy
}

func (geo *BufferGeometry) SetDrawPolicy(drawPolicy DrawPolicy) {
	geo.drawPolicy = drawPolicy
}

func (geo *BufferGeometry) NeedsUpdate() bool {
	return !geo.notNeedsUpdate
}

func (geo *BufferGeometry) SetNeedsUpdate(needsUpdate bool) {
	geo.notNeedsUpdate = !needsUpdate
}
