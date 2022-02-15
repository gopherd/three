package geometry

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
