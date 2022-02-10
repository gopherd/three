package object

import (
	"sync/atomic"

	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/math"
)

var nextObjectUUID int64

type Object interface {
	isObject()

	Type() string
	UUID() int64
	Tag() string
	SetTag(tag string)
	Parent() Object
	SetParent(parent Object)

	AddChild(child Object)
	RemoveChild(child Object)
	RemoveChildByIndex(i int)
	RemoveChildByTag(tag string)
	RemoveChildByUUID(uuid int64)

	NumChild() int
	GetChildByIndex(i int) Object
	GetChildByTag(tag string) Object
	GetChildByUUID(uuid int64) Object

	Transform() math.Mat4x4
	Render(renderer.Renderer, math.Mat4x4)
}

type Object3d struct {
	uuid      int64
	tag       string
	transform math.Mat4x4

	parent   Object
	children []Object
	byUUID   map[int64]int
	byTag    map[string]int
}

func (*Object3d) isObject() {}

func (obj *Object3d) Init() {
	obj.uuid = atomic.AddInt64(&nextObjectUUID)
}

func (obj *Object3d) Type() string            { return "" }
func (obj *Object3d) UUID() int64             { return obj.uuid }
func (obj *Object3d) Tag() string             { return obj.tag }
func (obj *Object3d) SetTag(tag string)       { obj.tag = tag }
func (obj *Object3d) Parent() Object          { return obj.parent }
func (obj *Object3d) SetParent(parent Object) { obj.parent = parent }

func (obj *Object3d) lazyInitByUUID() {
	if obj.byUUID == nil {
		obj.byUUID = make(map[int64]int)
	}
}

func (obj *Object3d) lazyInitByTag() {
	if obj.byTag == nil {
		obj.byTag = make(map[string]int)
	}
}

func (obj *Object3d) AddChild(child Object) {
	child.SetParent(obj)
	obj.lazyInitByUUID()
	if _, ok := obj.byUUID[child.UUID()]; ok {
		return
	}
	index := len(obj.children)
	obj.byUUID[child.UUID()] = index
	if tag := child.Tag(); tag != "" {
		obj.lazyInitByTag()
		obj.byTag[tag] = index
	}
	obj.children = append(obj.children, child)
}

func (obj *Object3d) RemoveChild(child Object) bool {
	child.SetParent(nil)
	if obj.byUUID == nil {
		return false
	}
	index, ok := obj.byUUID[child.UUID()]
	if !ok {
		return false
	}
}

func (obj *Object3d) RemoveChildByIndex(i int) {
}

func (obj *Object3d) RemoveChildByTag(tag string) {
}

func (obj *Object3d) RemoveChildByUUID(uuid int64) {
}

func (obj *Object3d) NumChild() int
func (obj *Object3d) GetChildByIndex(i int) Object
func (obj *Object3d) GetChildByTag(tag string) Object
func (obj *Object3d) GetChildByUUID(uuid int64) Object
