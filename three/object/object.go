package object

import (
	"sync/atomic"

	"github.com/gopherd/threego/three/driver/renderer"
	"github.com/gopherd/threego/three/math"
)

var nextObjectUUID int64

// Object 定义对象接口
type Object interface {
	isObject()
	setParent(parent Object)
	addChild(child Object) bool

	Type() string      // Type returns type of Object
	UUID() int64       // UUID returns UUID of Object
	Tag() string       // Tag returns tag of Object
	SetTag(tag string) // SetTag sets tag of Object
	Parent() Object    // Parent returns parent of Object

	RemoveChild(child Object) bool     // RemoveChild removes child object
	RemoveChildByIndex(i int)          // RemoveChildByIndex removes ith child object
	RemoveChildByTag(tag string) bool  // RemoveChildByTag removes child object by tag
	RemoveChildByUUID(uuid int64) bool // RemoveChildByUUID removes child object by uuid

	NumChild() int                    // NumChild returns number of children
	GetChildByIndex(i int) Object     // GetChildByIndex returns child object by index
	GetChildByTag(tag string) Object  // GetChildByTag retrieves child object by tag
	GetChildByUUID(uuid int64) Object // GetChildByUUID retrieves child object by uuid

	Transform() math.Mat4x4                                           // Transform returns transform matrix
	Render(renderer renderer.Renderer, camera, transform math.Mat4x4) // Render renders the Object to `renderer' with specified tranform
}

type node struct {
	children []Object
	byUUID   map[int64]int
	byTag    map[string]int
}

// addChild implements Object unexported addChild method
func (node *node) addChild(child Object) bool {
	if child.Parent() != nil {
		return false
	}
	if node.byUUID == nil {
		node.byUUID = make(map[int64]int)
	}
	var uuid = child.UUID()
	if _, ok := node.byUUID[uuid]; ok {
		return false
	}
	var index = len(node.children)
	node.byUUID[uuid] = index
	if tag := child.Tag(); tag != "" {
		if node.byTag == nil {
			node.byTag = make(map[string]int)
		}
		node.byTag[tag] = index
	}
	node.children = append(node.children, child)
	return true
}

func (node *node) removeChild(i int, child Object) {
	child.setParent(nil)
	delete(node.byUUID, child.UUID())
	if node.byTag != nil {
		if tag := child.Tag(); tag != "" {
			delete(node.byTag, tag)
		}
	}
	var end = len(node.children) - 1
	if i != end {
		node.children[i] = node.children[end]
		node.byUUID[node.children[i].UUID()] = i
		if node.byTag != nil {
			if tag := node.children[i].Tag(); tag != "" {
				node.byTag[tag] = i
			}
		}
	}
	node.children = node.children[:end]
}

// RemoveChild implements Object RemoveChild method
func (node *node) RemoveChild(child Object) bool {
	return node.RemoveChildByUUID(child.UUID())
}

// RemoveChildByIndex implements Object RemoveChildByIndex method
func (node *node) RemoveChildByIndex(i int) {
	node.removeChild(i, node.children[i])
}

// RemoveChildByTag implements Object RemoveChildByTag method
func (node *node) RemoveChildByTag(tag string) bool {
	if node.byTag == nil || tag == "" {
		return false
	}
	index, ok := node.byTag[tag]
	if !ok {
		return false
	}
	node.RemoveChildByIndex(index)
	return true
}

// RemoveChildByUUID implements Object RemoveChildByUUID method
func (node *node) RemoveChildByUUID(uuid int64) bool {
	if node.byUUID == nil {
		return false
	}
	index, ok := node.byUUID[uuid]
	if !ok {
		return false
	}
	node.RemoveChildByIndex(index)
	return true
}

// NumChild implements Object NumChild method
func (node *node) NumChild() int {
	return len(node.children)
}

// GetChildByIndex implements Object GetChildByIndex method
func (node *node) GetChildByIndex(i int) Object {
	return node.children[i]
}

// GetChildByTag implements Object GetChildByTag method
func (node *node) GetChildByTag(tag string) Object {
	if node.byTag == nil || tag == "" {
		return nil
	}
	index, ok := node.byTag[tag]
	if !ok {
		return nil
	}
	return node.children[index]
}

// GetChildByUUID implements Object GetChildByUUID method
func (node *node) GetChildByUUID(uuid int64) Object {
	if node.byUUID == nil {
		return nil
	}
	index, ok := node.byUUID[uuid]
	if !ok {
		return nil
	}
	return node.children[index]
}

type object3d struct {
	node
	parent  Object
	uuid    int64
	tag     string
	program struct {
		fail    bool
		created bool
		id      uint32
		vshader string
		fshader string
	}
	transform math.Mat4x4
}

// Init initializes Object
func (obj *object3d) Init() {
	obj.uuid = atomic.AddInt64(&nextObjectUUID, 1)
}

// isObject implements Object unexported isObject method
func (*object3d) isObject() {}

// Type implements Object Type method
func (obj *object3d) Type() string { return "" }

// UUID implements Object UUID method
func (obj *object3d) UUID() int64 { return obj.uuid }

// Tag implements Object Tag method
func (obj *object3d) Tag() string { return obj.tag }

// SetTag implements Object SetTag method
func (obj *object3d) SetTag(tag string) { obj.tag = tag }

// Parent implements Object Parent method
func (obj *object3d) Parent() Object { return obj.parent }

// setParent implements Object unexported setParent method
func (obj *object3d) setParent(parent Object) {
	obj.parent = parent
}

// Transform implements Object Transform method
func (obj *object3d) Transform() math.Mat4x4 {
	return obj.transform
}

func (obj *object3d) lazyInitProgram(renderer renderer.Renderer) error {
	if obj.program.created || obj.program.fail {
		return nil
	}
	program, err := renderer.CreateProgram(obj.program.vshader, obj.program.fshader)
	if err != nil {
		obj.program.fail = true
		return err
	}
	obj.program.created = true
	obj.program.id = program
	return nil
}

// Render implements Object Render method
func (obj *object3d) Render(renderer renderer.Renderer, cameraTransform, transform math.Mat4x4) {
	if err := obj.lazyInitProgram(renderer); err != nil {
		println(err.Error())
	}
	if obj.program.fail {
		return
	}
	renderer.SetUniform(obj.program.id, "view", cameraTransform)
	renderer.SetUniform(obj.program.id, "transform", transform)
}

// Add adds object child to parent
func Add(parent, child Object) {
	if parent.addChild(child) {
		child.setParent(parent)
	}
}

// MustAdd adds object child to parent. It would panic if failed
func MustAdd(parent, child Object) {
	if !parent.addChild(child) {
		panic("add child failed")
	}
	child.setParent(parent)
}
