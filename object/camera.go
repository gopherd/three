package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/driver/renderer"
	"github.com/gopherd/three/geometry"
)

type CameraType int

const (
	PerspectiveCameraType CameraType = iota
	OrthographicCameraType
)

// Camera represents a camera object
type Camera interface {
	Object

	CameraType() CameraType
	Projection() core.Matrix4
	SetViewOffset(fullWidth, fullHeight, x, y, width, height core.Float)

	IntersectsBox(box geometry.Box3) bool
	ContainsPoint(pos core.Vector3) bool
}

type cameraImpl struct {
	object3d
	matrixWorldInverse core.Matrix4
	proj               struct {
		matrix        core.Matrix4
		matrixInverse core.Matrix4
		frustum       geometry.Frustum
		view          struct {
			enabled               bool
			fullWidth, fullHeight core.Float
			offsetX, offsetY      core.Float
			width, height         core.Float
		}
		notNeedsUpdate bool
	}
	zoom      core.Float
	near, far core.Float
}

// TODO(delay) Bounds implements Object Bounds method
func (camera *cameraImpl) Bounds() geometry.Box3 {
	return geometry.Box3{}
}

// TODO(delay) Render implements Object Render method
func (camera *cameraImpl) Render(renderer renderer.Renderer, proj, view, transform core.Matrix4) {
}

// SetViewOffset implements Camera SetViewOffset method
func (camera *cameraImpl) SetViewOffset(fullWidth, fullHeight, x, y, width, height core.Float) {
	camera.proj.view.enabled = true
	camera.proj.view.fullWidth = fullWidth
	camera.proj.view.fullHeight = fullHeight
	camera.proj.view.offsetX = x
	camera.proj.view.offsetY = y
	camera.proj.view.width = width
	camera.proj.view.height = height
	camera.setProjectionNeedsUpdate(true)
}

// IntersectsBox implements Camera IntersectsBox method
func (camera *cameraImpl) IntersectsBox(box geometry.Box3) bool {
	return camera.proj.frustum.IntersectsBox(box)
}

// ContainsPoint implements Camera ContainsPoint method
func (camera *cameraImpl) ContainsPoint(point core.Vector3) bool {
	return camera.proj.frustum.ContainsPoint(point)
}

func (camera *cameraImpl) isProjectionNeedsUpdate() bool {
	return !camera.proj.notNeedsUpdate
}

func (camera *cameraImpl) setProjectionNeedsUpdate(needsUpdate bool) {
	camera.proj.notNeedsUpdate = !needsUpdate
}

func (camera *cameraImpl) projectionMatrixChanged() {
	camera.setProjectionNeedsUpdate(false)
	camera.proj.matrixInverse = camera.proj.matrix.Invert()
	camera.proj.frustum.SetFromProjectionMatrix(camera.proj.matrix)
}
