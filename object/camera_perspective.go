package object

import (
	"math"

	"github.com/gopherd/doge/math/mathutil"
	"github.com/gopherd/three/core"
)

// PerspectiveCamera represents a perspective camera
type PerspectiveCamera struct {
	cameraImpl

	fov, aspect core.Float
	filmGauge   core.Float // width of the film (default in millimeters)
	filmOffset  core.Float // horizontal film offset (same unit as gauge)
}

var _ Camera = (*PerspectiveCamera)(nil)

// NewPerspectiveCamera creates a PerspectiveCamera
func NewPerspectiveCamera(fov, aspect, near, far core.Float) *PerspectiveCamera {
	camera := new(PerspectiveCamera)
	camera.Init()
	camera.zoom = 1
	camera.fov = fov
	camera.aspect = aspect
	camera.near = near
	camera.far = far
	camera.updateProjectionMatrix()
	return camera
}

func (camera *PerspectiveCamera) String() string {
	return "PerspectiveCamera"
}

// CameraType implements Camera CameraType method
func (camera *PerspectiveCamera) CameraType() CameraType {
	return PerspectiveCameraType
}

// Projection implements Camera Projection method
func (camera *PerspectiveCamera) Projection() core.Matrix4 {
	camera.tryUpdateProjectionMatrix()
	return camera.proj.matrix
}

func (camera *PerspectiveCamera) tryUpdateProjectionMatrix() {
	if camera.isProjectionNeedsUpdate() {
		camera.updateProjectionMatrix()
	}
}

func (camera *PerspectiveCamera) updateProjectionMatrix() {
	camera.setProjectionNeedsUpdate(false)
	var near = camera.near
	var top = near * core.Float(math.Tan(mathutil.Deg2Rad(0.5*float64(camera.fov)))/float64(camera.zoom))
	var height = 2 * top
	var width = camera.aspect * height
	var left = -0.5 * width
	if camera.proj.view.enabled {
		var fullWidth = camera.proj.view.fullWidth
		var fullHeight = camera.proj.view.fullHeight
		left += camera.proj.view.offsetX * width / fullWidth
		top -= camera.proj.view.offsetY * height / fullHeight
		width *= camera.proj.view.width / fullWidth
		height *= camera.proj.view.height / fullHeight
	}

	var skew = camera.filmOffset
	if skew != 0 {
		left += near * skew / camera.GetFilmWidth()
	}

	camera.proj.matrix.MakePerspective(left, left+width, top, top-height, near, camera.far)
	camera.projectionMatrixChanged()
}

// see {@link http://www.bobatkins.com/photography/technical/field_of_view.html}
func (camera *PerspectiveCamera) SetFocalLength(focalLength core.Float) {
	var vExtentSlope = 0.5 * camera.GetFilmHeight() / focalLength
	camera.fov = core.Float(mathutil.Rad2Deg(2 * math.Atan(float64(vExtentSlope))))
	camera.setProjectionNeedsUpdate(true)
}

func (camera *PerspectiveCamera) GetFocalLength() core.Float {
	var vExtentSlope = core.Float(math.Tan(float64(mathutil.Deg2Rad(0.5 * camera.fov))))
	return 0.5 * camera.GetFilmHeight() / vExtentSlope
}

func (camera *PerspectiveCamera) GetEffectiveFOV() core.Float {
	return core.Float(mathutil.Rad2Deg(2 * math.Atan(math.Tan(mathutil.Deg2Rad(0.5*float64(camera.fov))/float64(camera.zoom)))))
}

func (camera *PerspectiveCamera) GetFilmWidth() core.Float {
	// film not completely covered in portrait format (aspect < 1)
	return camera.filmGauge * mathutil.Min(camera.aspect, 1)
}

func (camera *PerspectiveCamera) GetFilmHeight() core.Float {
	// film not completely covered in landscape format (aspect > 1)
	return camera.filmGauge / mathutil.Max(camera.aspect, 1)
}
