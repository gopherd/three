package math

type Mat4x4 [4 * 4]float32

var One4x4 = Mat4x4{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

func (mat Mat4x4) Mul(mat2 Mat4x4) Mat4x4 {
	var result Mat4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			index := i + j*4
			for k := 0; k < 4; k++ {
				result[index] += mat[k+j*4] * mat2[i+k*4]
			}
		}
	}
	return result
}
