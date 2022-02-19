package object

import (
	"github.com/gopherd/three/core"
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

func (camera *OrthographicCamera) ToString() string {
	return "OrthographicCamera.ToString:TODO"
}

// CameraType implements Camera CameraType method
func (camera *OrthographicCamera) CameraType() CameraType {
	return OrthographicCameraType
}

// Projection implements Object Projection method
func (camera *OrthographicCamera) Projection() core.Matrix4 {
	camera.tryUpdateProjectionMatrix()
	return camera.proj.matrix
}

func (camera *OrthographicCamera) tryUpdateProjectionMatrix() {
	if camera.isProjectionNeedsUpdate() {
		camera.updateProjectionMatrix()
	}
}

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
	camera.projectionMatrixChanged()
}
