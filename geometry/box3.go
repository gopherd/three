package geometry

import (
	"github.com/gopherd/doge/math/mathutil"

	"github.com/gopherd/three/core"
)

type Box3 struct {
	Min, Max core.Vector3
}

func (box Box3) Center() core.Vector3 { return box.Min.Add(box.Max).Div(2) }
func (box Box3) Size() core.Vector3   { return box.Max.Sub(box.Min) }

func (box Box3) IsEmpty() bool {
	return box.Max.X() < box.Min.X() || box.Max.Y() < box.Min.Y() || box.Max.Z() < box.Min.Z()
}

func (box Box3) ContainsPoint(point core.Vector3) bool {

	return !(point.X() < box.Min.X() || point.X() > box.Max.X() ||
		point.Y() < box.Min.Y() || point.Y() > box.Max.Y() ||
		point.Z() < box.Min.Z() || point.Z() > box.Max.Z())

}

func (box Box3) ContainsBox(other Box3) bool {
	return box.Min.X() <= other.Min.X() && other.Max.X() <= box.Max.X() &&
		box.Min.Y() <= other.Min.Y() && other.Max.Y() <= box.Max.Y() &&
		box.Min.Z() <= other.Min.Z() && other.Max.Z() <= box.Max.Z()
}

func (box Box3) Intersect(other Box3) Box3 {
	var min = core.Vec3(
		mathutil.Max(box.Min.X(), other.Min.X()),
		mathutil.Max(box.Min.Y(), other.Min.Y()),
		mathutil.Max(box.Min.Z(), other.Min.Z()),
	)
	var max = core.Vec3(
		mathutil.Min(box.Max.X(), other.Max.X()),
		mathutil.Min(box.Max.Y(), other.Max.Y()),
		mathutil.Min(box.Max.Z(), other.Max.Z()),
	)
	return Box3{Min: min, Max: max}
}

func (box Box3) Union(other Box3) Box3 {
	var min = core.Vec3(
		mathutil.Min(box.Min.X(), other.Min.X()),
		mathutil.Min(box.Min.Y(), other.Min.Y()),
		mathutil.Min(box.Min.Z(), other.Min.Z()),
	)
	var max = core.Vec3(
		mathutil.Max(box.Max.X(), other.Max.X()),
		mathutil.Max(box.Max.Y(), other.Max.Y()),
		mathutil.Max(box.Max.Z(), other.Max.Z()),
	)
	return Box3{Min: min, Max: max}
}

func (box Box3) IntersectsBox(other Box3) bool {
	return !(other.Max.X() < box.Min.X() || other.Min.X() > box.Max.X() ||
		other.Max.Y() < box.Min.Y() || other.Min.Y() > box.Max.Y() ||
		other.Max.Z() < box.Min.Z() || other.Min.Z() > box.Max.Z())

}

func (box Box3) IntersectsSphere(sphere Sphere3) bool {
	// Find the point on the AABB closest to the sphere center.
	var p = box.ClampPoint(sphere.Center)
	return p.Sub(sphere.Center).Square() <= sphere.Radius*sphere.Radius
}

func (box Box3) IntersectsPlane(plane Plane) bool {
	var min, max core.Float
	if plane.Normal.X() > 0 {
		min = plane.Normal.X() * box.Min.X()
		max = plane.Normal.X() * box.Max.X()
	} else {
		min = plane.Normal.X() * box.Max.X()
		max = plane.Normal.X() * box.Min.X()
	}
	if plane.Normal.Y() > 0 {
		min += plane.Normal.Y() * box.Min.Y()
		max += plane.Normal.Y() * box.Max.Y()
	} else {
		min += plane.Normal.Y() * box.Max.Y()
		max += plane.Normal.Y() * box.Min.Y()
	}
	if plane.Normal.Z() > 0 {
		min += plane.Normal.Z() * box.Min.Z()
		max += plane.Normal.Z() * box.Max.Z()
	} else {
		min += plane.Normal.Z() * box.Max.Z()
		max += plane.Normal.Z() * box.Min.Z()
	}
	return min <= -plane.Constant && max >= -plane.Constant
}

func (box Box3) ClampPoint(point core.Vector3) core.Vector3 {
	return core.Vec3(
		mathutil.Clamp(point.X(), box.Min.X(), box.Max.X()),
		mathutil.Clamp(point.Y(), box.Min.Y(), box.Max.Y()),
		mathutil.Clamp(point.Z(), box.Min.Z(), box.Max.Z()),
	)
}

func (box Box3) DistanceToPoint(point core.Vector3) core.Float {
	return box.ClampPoint(point).Sub(point).Length()
}
