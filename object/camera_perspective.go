package object

import (
	"math"

	"github.com/gopherd/doge/math/mathutil"
	"github.com/gopherd/three/core"
	"github.com/gopherd/three/geometry"
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

// IsPerspectiveCamera implements Camera IsPerspectiveCamera method
func (camera *PerspectiveCamera) CameraType() CameraType {
	return PerspectiveCameraType
}

// TODO(delay) IntersectsBox implements Camera IntersectsBox method
func (camera *PerspectiveCamera) IntersectsBox(box geometry.Box3) bool {
	return true
}

// TODO(delay) ContainsPoint implements Camera ContainsPoint method
func (camera *PerspectiveCamera) ContainsPoint(point core.Vector3) bool {
	return true
}

// Projection implements Camera Projection method
func (camera *PerspectiveCamera) Projection() core.Matrix4 {
	if camera.isProjectionNeedsUpdate() {
		camera.updateProjectionMatrix()
	}
	return camera.proj.matrix
}

// updateProjectionMatrix try updates projection matrix
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
	camera.proj.matrixInverse = camera.proj.matrix.Invert()
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
