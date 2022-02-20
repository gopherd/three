package object

import (
	"bytes"
	"sync/atomic"

	"github.com/gopherd/doge/container"
	"github.com/gopherd/doge/container/stringify"
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/core/event"
	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/driver/renderer/shader"
	"github.com/gopherd/three/geometry"
	"github.com/gopherd/three/material"
)

var nextObjectUUID int64

// Object reprensents object in scene
type Object interface {
	node

	Type() string      // Type returns type of Object
	UUID() int64       // UUID returns UUID of Object
	Tag() string       // Tag returns tag of Object
	SetTag(tag string) // SetTag sets tag of Object

	Visible() bool                          // Visible reports whether the object is visible
	Bounds() geometry.Box3                  // Bounds returns object bounding box
	Transform() core.Matrix4                // Transform returns transform matrix in local space
	TransformWorld() core.Matrix4           // TransformWorld returns transform matrix in world space
	LocalToWorld(core.Vector3) core.Vector3 // LocalToWorld Converts the vector from this object's local space to world space
	LookAt(core.Vector3)                    // LookAt looks at a position in world space

	// Render renders the Object to `renderer' with specified matrices
	Render(renderer renderer.Renderer, proj, view, transform core.Matrix4)
}

type node interface {
	container.Node[Object]

	addChild(child Object)
	setParent(parent Object)

	RemoveChild(child Object) bool     // RemoveChild removes child object
	RemoveChildByIndex(i int)          // RemoveChildByIndex removes ith child object
	RemoveChildByTag(tag string) bool  // RemoveChildByTag removes child object by tag
	RemoveChildByUUID(uuid int64) bool // RemoveChildByUUID removes child object by uuid

	GetChildByTag(tag string) Object  // GetChildByTag retrieves child object by tag
	GetChildByUUID(uuid int64) Object // GetChildByUUID retrieves child object by uuid

	DispatchEvent(event.Event) bool
	OnUpdate()
}

type node3d struct {
	event.Dispatcher

	parent   Object
	children []Object
	byUUID   map[int64]int
	byTag    map[string]int
}

// Parent implements node container.Node Parent method
func (node *node3d) Parent() Object { return node.parent }

// NumChild implements container.Node NumChild method
func (node *node3d) NumChild() int {
	return len(node.children)
}

// GetChildByIndex implements container.Node GetChildByIndex method
func (node *node3d) GetChildByIndex(i int) Object {
	return node.children[i]
}

// setParent implements Object unexported setParent method
func (node *node3d) setParent(parent Object) {
	node.parent = parent
}

// addChild implements Object unexported addChild method
func (node *node3d) addChild(child Object) {
	if parent := child.Parent(); parent != nil {
		parent.RemoveChild(child)
	}
	if node.byUUID == nil {
		node.byUUID = make(map[int64]int)
	}
	var uuid = child.UUID()
	if _, ok := node.byUUID[uuid]; ok {
		return
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
	child.DispatchEvent(addedEvent)
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
	node.children[end] = nil
	node.children = node.children[:end]
	child.DispatchEvent(removedEvent)
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
	uuid    int64
	tag     string
	program struct {
		renderer.Program
		created bool
		fail    bool
	}
	invisible bool
	transform struct {
		position       core.Vector3
		scale          core.Vector3
		rotation       core.Euler
		quaternion     core.Vector4
		matrix         core.Matrix4
		notNeedsUpdate bool
	}
	transformWorld struct {
		matrix         core.Matrix4
		notNeedsUpdate bool
	}
}

// Init initializes Object
func (obj *object3d) Init() {
	obj.uuid = atomic.AddInt64(&nextObjectUUID, 1)
	obj.transform.matrix.MakeIdentity()
	obj.transformWorld.matrix.MakeIdentity()
	// TODO: update position, scale, rotation, quaternion
}

func (obj *object3d) String() string {
	var buf bytes.Buffer
	// TODO: write obect information
	return buf.String()
}

// Type implements Object Type method
func (obj *object3d) Type() string { return "" }

// UUID implements Object UUID method
func (obj *object3d) UUID() int64 { return obj.uuid }

// Tag implements Object Tag method
func (obj *object3d) Tag() string { return obj.tag }

// SetTag implements Object SetTag method
func (obj *object3d) SetTag(tag string) { obj.tag = tag }

// Visible implements Object Visible method
func (obj *object3d) Visible() bool {
	return !obj.invisible
}

// SetVisible sets object visible property
func (obj *object3d) SetVisible(visible bool) {
	obj.invisible = !visible
}

// Transform implements Object Transform method
func (obj *object3d) Transform() core.Matrix4 {
	return obj.transform.matrix
}

// TransformWorld implements Object TransformWorld method
func (obj *object3d) TransformWorld() core.Matrix4 {
	if obj.parent == nil {
		return obj.transform.matrix
	}
	return obj.parent.TransformWorld().Dot(obj.transform.matrix)
}

func (obj object3d) GetPosition() core.Vector3 {
	return obj.transform.position
}

func (obj *object3d) SetPosition(pos core.Vector3) {
	obj.transform.position = pos
	obj.transform.notNeedsUpdate = false
}

func (obj object3d) GetScale() core.Vector3 {
	return obj.transform.scale
}

func (obj *object3d) SetScale(scale core.Vector3) {
	obj.transform.scale = scale
	obj.transform.notNeedsUpdate = false
}

func (obj object3d) GetRotation() core.Euler {
	return obj.transform.rotation
}

func (obj *object3d) SetRotation(euler core.Euler) {
	obj.transform.rotation = euler
	obj.transform.notNeedsUpdate = false
}

func (obj object3d) GetQuaternion() core.Vector4 {
	return obj.transform.quaternion
}

func (obj *object3d) SetQuaternion(quaternion core.Vector4) {
	obj.transform.quaternion = quaternion
	obj.transform.notNeedsUpdate = false
}

// LocalToWorld implements Object LocalToWorld method
func (obj *object3d) LocalToWorld(vec core.Vector3) core.Vector3 {
	return obj.TransformWorld().DotVec3(vec)
}

// TODO: LookAt implements Object LookAt method
func (obj *object3d) LookAt(pos core.Vector3) {
}

func (obj *object3d) createProgram(renderer renderer.Renderer, shader shader.Shader) error {
	program, err := renderer.CreateProgram(shader.Vertex, shader.Fragment)
	if err != nil {
		obj.program.fail = false
		return err
	}
	obj.program.created = true
	obj.program.Program = program
	return nil
}

// Render implements Object Render method
func (obj *object3d) Render(renderer renderer.Renderer, proj, view, transform core.Matrix4) {
	renderer.SetUniform(obj.program.Id, "proj", proj)
	renderer.SetUniform(obj.program.Id, "view", view)
	renderer.SetUniform(obj.program.Id, "transform", transform)
}

func (obj *object3d) renderGeometry(
	renderer renderer.Renderer,
	geometry geometry.Geometry,
	material material.Material,
) {
	var shader = material.Shader()
	if !obj.program.created && !obj.program.fail {
		if err := obj.createProgram(renderer, shader); err != nil {
			panic(err)
		}
	}
	var needsUpdate = material.NeedsUpdate()
	if needsUpdate {
		material.SetNeedsUpdate(false)
		for name, uniform := range shader.Uniforms {
			renderer.SetUniform(obj.program.Id, name, uniform)
		}
	}

	var attributes = geometry.Attributes()
	var index = geometry.Index()
	if geometry.NeedsUpdate() {
		geometry.SetNeedsUpdate(false)
		// TODO: update attributes
	}
	_, _ = attributes, index
}

// Attatch attatchs child to parent object
func Attatch(parent, child Object) {
	parent.addChild(child)
	child.setParent(parent)
}

func recursivelyRenderObject(
	renderer renderer.Renderer,
	camera Camera,
	proj, view core.Matrix4,
	object Object,
	transform core.Matrix4,
) {
	renderObject(renderer, camera, proj, view, object, transform)
	for i, n := 0, object.NumChild(); i < n; i++ {
		child := object.GetChildByIndex(i)
		if !child.Visible() {
			continue
		}
		childTransform := transform.Dot(child.Transform())
		recursivelyRenderObject(renderer, camera, proj, view, child, childTransform)
	}
}

func renderObject(
	renderer renderer.Renderer,
	camera Camera,
	proj, view core.Matrix4,
	object Object,
	transform core.Matrix4,
) {
	box := object.Bounds()
	if !box.IsEmpty() {
		box.Min = transform.DotVec3(box.Min)
		box.Max = transform.DotVec3(box.Max)
		if !camera.IntersectsBox(box) {
			return
		}
	}
	object.Render(renderer, proj, view, transform)
}

func recursivelyUpdateNode(node node) {
	node.OnUpdate()
	for i, n := 0, node.NumChild(); i < n; i++ {
		recursivelyUpdateNode(node.GetChildByIndex(i))
	}
}

func Stringify(node node, options *stringify.Options) string {
	return stringify.Stringify[Object](node, options)
}
