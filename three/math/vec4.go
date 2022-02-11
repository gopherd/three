package math

import (
	"image/color"
	"math"
)

// Vector4 implements 4d vector
type Vector4 [4]float32

func Vec4(x, y, z, w float32) Vector4 {
	return Vector4{x, y, z, w}
}

func Color(c color.Color) Vector4 {
	r, g, b, a := c.RGBA()
	return Vec4(
		float32(r)/0xffff,
		float32(g)/0xffff,
		float32(b)/0xffff,
		float32(a)/0xffff,
	)
}

func (vec Vector4) X() float32 { return vec[0] }
func (vec Vector4) Y() float32 { return vec[1] }
func (vec Vector4) Z() float32 { return vec[2] }
func (vec Vector4) W() float32 { return vec[3] }

func (vec Vector4) R() float32 { return vec[0] }
func (vec Vector4) G() float32 { return vec[1] }
func (vec Vector4) B() float32 { return vec[2] }
func (vec Vector4) A() float32 { return vec[3] }

func (vec Vector4) Vec3() Vector3 {
	if vec[3] == 0 {
		return Vec3(vec[0], vec[1], vec[2])
	}
	return Vec3(vec[0]/vec[3], vec[1]/vec[3], vec[2]/vec[3])
}

func (vec Vector4) Sum() float32 {
	return vec[0] + vec[1] + vec[2] + vec[3]
}

func (vec Vector4) Dot(other Vector4) float32 {
	return vec[0]*other[0] + vec[1]*other[1] + vec[2]*other[2] + vec[3]*other[3]
}

func (vec Vector4) Square() float32 {
	return vec.Dot(vec)
}

func (vec Vector4) Length() float32 {
	return float32(math.Sqrt(float64(vec.Square())))
}

func (vec Vector4) Add(other Vector4) Vector4 {
	return Vec4(vec[0]+other[0], vec[1]+other[1], vec[2]+other[2], vec[3]+other[3])
}

func (vec Vector4) Sub(other Vector4) Vector4 {
	return Vec4(vec[0]-other[0], vec[1]-other[1], vec[2]-other[2], vec[3]-other[3])
}

func (vec Vector4) Mul(k float32) Vector4 {
	return Vec4(vec[0]*k, vec[1]*k, vec[2]*k, vec[3]*k)
}

func (vec Vector4) Div(k float32) Vector4 {
	return Vec4(vec[0]/k, vec[1]/k, vec[2]/k, vec[3]/k)
}

func (vec Vector4) Hadamard(other Vector4) Vector4 {
	return Vec4(vec[0]*other[0], vec[1]*vec[1], vec[2]*other[2], vec[3]*other[3])
}
