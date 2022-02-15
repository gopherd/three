package core

import (
	"image/color"

	"github.com/gopherd/doge/math/tensor"
)

type Float = float32

type Vector2 = tensor.Vector2[Float]
type Vector3 = tensor.Vector3[Float]
type Vector4 = tensor.Vector4[Float]
type Matrix4 = tensor.Matrix4[Float]
type Euler = tensor.Euler[Float]

func Vec2(x, y Float) Vector2       { return tensor.Vec2(x, y) }
func Vec3(x, y, z Float) Vector3    { return tensor.Vec3(x, y, z) }
func Vec4(x, y, z, w Float) Vector4 { return tensor.Vec4(x, y, z, w) }
func One4x4() Matrix4               { return tensor.One4x4[Float]() }

func Color(c color.Color) Vector4 {
	const max = 0xffff
	var r, g, b, a = c.RGBA()
	return tensor.Vec4(Float(r)/max, Float(g)/max, Float(b)/max, Float(a)/max)
}
