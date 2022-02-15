package object

import (
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/geometry"
)

// OrthographicCamera represents an orthographic camera
type OrthographicCamera struct {
	cameraImpl

	left, right, top, bottom core.Float
}

var _ Camera = (*OrthographicCamera)(nil)

// NewOrthographicCamera creates a OrthographicCamera
func NewOrthographicCamera(left, right, top, bottom, near, far core.Float) *OrthographicCamera {
	camera := new(OrthographicCamera)
	camera.Init()
	camera.zoom = 1
	camera.left = left
	camera.right = right
	camera.top = top
	camera.bottom = bottom
	camera.near = near
	camera.far = far
	camera.updateProjectionMatrix()
	return camera
}

// CameraType implements Camera CameraType method
func (camera *OrthographicCamera) CameraType() CameraType {
	return OrthographicCameraType
}

// Projection implements Object Projection method
func (camera *OrthographicCamera) Projection() core.Matrix4 {
	if camera.isProjectionNeedsUpdate() {
		camera.setProjectionNeedsUpdate(false)
		camera.updateProjectionMatrix()
	}
	return camera.proj.matrix
}

// TODO(delay) IntersectsBox implements Camera IntersectsBox method
func (camera *OrthographicCamera) IntersectsBox(box geometry.Box3) bool {
	return true
}

// TODO(delay) ContainsPoint implements Camera ContainsPoint method
func (camera *OrthographicCamera) ContainsPoint(point core.Vector3) bool {
	return true
}

// updateProjectionMatrix try updates projection matrix
func (camera *OrthographicCamera) updateProjectionMatrix() {
	var dx = (camera.right - camera.left) / (2 * camera.zoom)
	var dy = (camera.top - camera.bottom) / (2 * camera.zoom)
	var cx = (camera.right + camera.left) / 2
	var cy = (camera.top + camera.bottom) / 2

	var left = cx - dx
	var right = cx + dx
	var top = cy + dy
	var bottom = cy - dy

	if camera.proj.view.enabled {
		var scaleW = (camera.right - camera.left) / camera.proj.view.fullWidth / camera.zoom
		var scaleH = (camera.top - camera.bottom) / camera.proj.view.fullHeight / camera.zoom
		left += scaleW * camera.proj.view.offsetX
		right = left + scaleW*camera.proj.view.width
		top -= scaleH * camera.proj.view.offsetY
		bottom = top - scaleH*camera.proj.view.height
	}

	camera.proj.matrix.MakeOrthographic(left, right, top, bottom, camera.near, camera.far)
	camera.proj.matrixInverse = camera.proj.matrix.Invert()
}

// TODO(delay) ContainsBox implements Camera ContainsBox method
func (camera *OrthographicCamera) ContainsBox(proj, view core.Matrix4, min, max core.Vector3) bool {
	return true
}
