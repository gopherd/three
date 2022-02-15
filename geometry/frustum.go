package geometry

import (
	"github.com/gopherd/doge/operator"

	"github.com/gopherd/three/core"
)

type Frustum [6]Plane

func (frustum *Frustum) SetFromProjectionMatrix(m core.Matrix4) *Frustum {
	var me0, me1, me2, me3 = m[0], m[1], m[2], m[3]
	var me4, me5, me6, me7 = m[4], m[5], m[6], m[7]
	var me8, me9, me10, me11 = m[8], m[9], m[10], m[11]
	var me12, me13, me14, me15 = m[12], m[13], m[14], m[15]

	(*frustum)[0].SetComponents(me3-me0, me7-me4, me11-me8, me15-me12).Normalize()
	(*frustum)[1].SetComponents(me3+me0, me7+me4, me11+me8, me15+me12).Normalize()
	(*frustum)[2].SetComponents(me3+me1, me7+me5, me11+me9, me15+me13).Normalize()
	(*frustum)[3].SetComponents(me3-me1, me7-me5, me11-me9, me15-me13).Normalize()
	(*frustum)[4].SetComponents(me3-me2, me7-me6, me11-me10, me15-me14).Normalize()
	(*frustum)[5].SetComponents(me3+me2, me7+me6, me11+me10, me15+me14).Normalize()

	return frustum
}

func (frustum Frustum) IntersectsSphere(sphere Sphere3) bool {
	var center = sphere.Center
	var negRadius = -sphere.Radius
	for i := range frustum {
		var distance = frustum[i].DistanceToPoint(center)
		if distance < negRadius {
			return false
		}
	}
	return true
}

func (frustum Frustum) IntersectsBox(box Box3) bool {
	for i := range frustum {
		var x = operator.Conditional(frustum[i].Normal.X() > 0, box.Max.X(), box.Min.X())
		var y = operator.Conditional(frustum[i].Normal.Y() > 0, box.Max.Y(), box.Min.Y())
		var z = operator.Conditional(frustum[i].Normal.Z() > 0, box.Max.Z(), box.Min.Z())
		if frustum[i].DistanceToPoint(core.Vec3(x, y, z)) < 0 {
			return false
		}
	}
	return true
}

func (frustum Frustum) ContainsPoint(point core.Vector3) bool {
	for i := range frustum {
		if frustum[i].DistanceToPoint(point) < 0 {
			return false
		}
	}
	return true
}
