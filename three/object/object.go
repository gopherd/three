package object

import (
	"sync/atomic"

	"github.com/gopherd/doge/math/tensor"

	"github.com/gopherd/threego/three/driver/renderer"
)

var nextObjectUUID int64

// Object reprensents object in scene
type Object interface {
	node
	setParent(parent Object)

	Type() string      // Type returns type of Object
	UUID() int64       // UUID returns UUID of Object
	Tag() string       // Tag returns tag of Object
	SetTag(tag string) // SetTag sets tag of Object
	Parent() Object    // Parent returns parent of Object

	Visible() bool                              // Visible reports whether the object is visible
	Bounds() (min, max tensor.Vector3, ok bool) // Bounds returns object bounding box
	Transform() tensor.Mat4x4                   // Transform returns transform matrix

	// Render renders the Object to `renderer' with specified tranform
	Render(renderer renderer.Renderer, cameraTransform, transform tensor.Mat4x4)
}

type node interface {
	addChild(child Object) bool

	RemoveChild(child Object) bool     // RemoveChild removes child object
	RemoveChildByIndex(i int)          // RemoveChildByIndex removes ith child object
	RemoveChildByTag(tag string) bool  // RemoveChildByTag removes child object by tag
	RemoveChildByUUID(uuid int64) bool // RemoveChildByUUID removes child object by uuid

	NumChild() int                    // NumChild returns number of children
	GetChildByIndex(i int) Object     // GetChildByIndex returns child object by index
	GetChildByTag(tag string) Object  // GetChildByTag retrieves child object by tag
	GetChildByUUID(uuid int64) Object // GetChildByUUID retrieves child object by uuid

	OnUpdate()
}

type node3d struct {
	children []Object
	byUUID   map[int64]int
	byTag    map[string]int

	onChildAdded   func(child Object)
	onChildRemoved func(child Object)
}

// addChild implements Object unexported addChild method
func (node *node3d) addChild(child Object) bool {
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

func (node *node3d) removeChild(i int, child Object) {
	parent := child.Parent()
	if parent != nil {
		child.setParent(nil)
	}
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
	if node.onChildRemoved != nil {
		node.onChildRemoved(child)
	}
}

// RemoveChild implements Object RemoveChild method
func (node *node3d) RemoveChild(child Object) bool {
	return node.RemoveChildByUUID(child.UUID())
}

// RemoveChildByIndex implements Object RemoveChildByIndex method
func (node *node3d) RemoveChildByIndex(i int) {
	node.removeChild(i, node.children[i])
}

// RemoveChildByTag implements Object RemoveChildByTag method
func (node *node3d) RemoveChildByTag(tag string) bool {
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
func (node *node3d) RemoveChildByUUID(uuid int64) bool {
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
func (node *node3d) NumChild() int {
	return len(node.children)
}

// GetChildByIndex implements Object GetChildByIndex method
func (node *node3d) GetChildByIndex(i int) Object {
	return node.children[i]
}

// GetChildByTag implements Object GetChildByTag method
func (node *node3d) GetChildByTag(tag string) Object {
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
func (node *node3d) GetChildByUUID(uuid int64) Object {
	if node.byUUID == nil {
		return nil
	}
	index, ok := node.byUUID[uuid]
	if !ok {
		return nil
	}
	return node.children[index]
}

// OnUpdate implements Object OnUpdate method
func (node *node3d) OnUpdate() {}

type object3d struct {
	node3d
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
	visible   bool
	transform tensor.Mat4x4
}

// Init initializes Object
func (obj *object3d) Init() {
	obj.uuid = atomic.AddInt64(&nextObjectUUID, 1)
}

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

// Visible implements Object Visible method
func (obj *object3d) Visible() bool {
	return obj.visible
}

// SetVisible sets object visible property
func (obj *object3d) SetVisible(visible bool) {
	obj.visible = visible
}

// Transform implements Object Transform method
func (obj *object3d) Transform() tensor.Mat4x4 {
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
func (obj *object3d) Render(renderer renderer.Renderer, camera, transform tensor.Mat4x4) {
	if err := obj.lazyInitProgram(renderer); err != nil {
		println(err.Error())
	}
	if obj.program.fail {
		return
	}
	renderer.SetUniform(obj.program.id, "view", camera)
	renderer.SetUniform(obj.program.id, "transform", transform)
}

// Attatch attatchs child to parent object
func Attatch(parent, child Object) bool {
	if parent.addChild(child) {
		child.setParent(parent)
		return true
	}
	return false
}

// TransformToWorld calculates object's transform in world
func TransformToWorld(obj Object) tensor.Mat4x4 {
	var mat = obj.Transform()
	var parent = obj.Parent()
	for parent != nil {
		mat = parent.Transform().Dot(mat)
		parent = parent.Parent()
	}
	return mat
}

func recursivelyRenderObject(
	renderer renderer.Renderer,
	camera Camera,
	cameraTransform tensor.Mat4x4,
	object Object,
	objectTransform tensor.Mat4x4,
) {
	renderObject(renderer, camera, cameraTransform, object, objectTransform)
	for i, n := 0, object.NumChild(); i < n; i++ {
		child := object.GetChildByIndex(i)
		if !child.Visible() {
			continue
		}
		childTransform := objectTransform.Dot(child.Transform())
		recursivelyRenderObject(renderer, camera, cameraTransform, child, childTransform)
	}
}

func renderObject(
	renderer renderer.Renderer,
	camera Camera,
	cameraTransform tensor.Mat4x4,
	object Object,
	objectTransform tensor.Mat4x4,
) {
	min, max, ok := object.Bounds()
	if ok {
		min = objectTransform.DotVec3(min)
		max = objectTransform.DotVec3(max)
		if !camera.ContainsBox(cameraTransform, min, max) {
			return
		}
	}
	object.Render(renderer, cameraTransform, objectTransform)
}

func recursivelyUpdateNode(node node) {
	node.OnUpdate()
	for i, n := 0, node.NumChild(); i < n; i++ {
		recursivelyUpdateNode(node.GetChildByIndex(i))
	}
}
