package math_test

import (
	"testing"

	"github.com/gopherd/threego/three/math"
)

func TestSum(t *testing.T) {
	type testCase struct {
		vec math.Vector2
		sum float32
	}
	for i, tc := range []testCase{
		{math.Vec2(0, 0), 0},
		{math.Vec2(1, 0), 1},
		{math.Vec2(0, 1), 1},
		{math.Vec2(1, 1), 2},
		{math.Vec2(0.5, 0.5), 1},
	} {
		sum := tc.vec.Sum()
		if sum != tc.sum {
			t.Fatalf("%dth: want %f, got %f", i, tc.sum, sum)
		}
	}
}

func TestAdd(t *testing.T) {
	type testCase struct {
		v1, v2, v3 math.Vector2
	}
	for i, tc := range []testCase{
		{math.Vec2(1, 2), math.Vec2(3, 4), math.Vec2(4, 6)},
	} {
		v := tc.v1.Add(tc.v2)
		if v != tc.v3 {
			t.Fatalf("%dth: want %v, got %v", i, tc.v3, v)
		}
	}
}
