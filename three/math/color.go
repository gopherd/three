package math

import "image/color"

func NormalizeColor(c color.Color) (r, g, b, a float32) {
	r0, g0, b0, a0 := c.RGBA()
	return float32(r0) / 0xffff, float32(g0) / 0xffff, float32(b0) / 0xffff, float32(a0) / 0xffff
}
