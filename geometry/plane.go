package geometry

import "github.com/gopherd/three/core"

type Plane struct {
	Normal   core.Vector3
	Constant core.Float
}

func (plane *Plane) Set(normal core.Vector3, constant core.Float) *Plane {
	plane.Normal = normal
	plane.Constant = constant
	return plane
}

func (plane *Plane) SetComponents(x, y, z, w core.Float) *Plane {
	plane.Normal.Set(x, y, z)
	plane.Constant = w
	return plane
}

func (plane *Plane) SetFromNormalAndCoplanarPoint(normal, point core.Vector3) *Plane {
	plane.Normal = normal
	plane.Constant = -point.Dot(normal)
	return plane
}

func (plane *Plane) SetFromCoplanarPoints(a, b, c core.Vector3) *Plane {
	var normal = c.Sub(b).Cross(a.Sub(b)).Normalize()
	plane.SetFromNormalAndCoplanarPoint(normal, a)
	return plane
}

func (plane *Plane) Normalize() *Plane {
	var inverseNormalLength = 1.0 / plane.Normal.Length()
	plane.Normal.Mul(inverseNormalLength)
	plane.Constant *= inverseNormalLength
	return plane
}

func (plane *Plane) Negate() *Plane {
	plane.Constant *= -1
	plane.Normal = plane.Normal.Mul(-1)
	return plane
}

func (plane Plane) DistanceToPoint(point core.Vector3) core.Float {
	return plane.Normal.Dot(point) + plane.Constant
}

func (plane Plane) DistanceToSphere(sphere Sphere3) core.Float {
	return plane.DistanceToPoint(sphere.Center) - sphere.Radius
}

func (plane Plane) ProjectPoint(point core.Vector3) core.Vector3 {
	return plane.Normal.Mul(-plane.DistanceToPoint(point)).Add(point)
}

func (plane Plane) IntersectLine(line Line3) (p core.Vector3, ok bool) {
	var direction = line.Direction()
	var denominator = plane.Normal.Dot(direction)
	if denominator == 0 {
		// line is coplanar, return origin
		if plane.DistanceToPoint(line.Start) == 0 {
			return line.Start, true
		}
		return

	}
	var t = -(line.Start.Dot(plane.Normal) + plane.Constant) / denominator
	if t < 0 || t > 1 {
		return
	}
	return direction.Mul(t).Add(line.Start), true
}

func (plane Plane) IntersectsLine(line Line3) bool {
	var startSign = plane.DistanceToPoint(line.Start)
	var endSign = plane.DistanceToPoint(line.End)
	return (startSign < 0 && endSign > 0) || (endSign < 0 && startSign > 0)

}

func (plane Plane) CoplanarPoint() core.Vector3 {
	return plane.Normal.Mul(-plane.Constant)
}
