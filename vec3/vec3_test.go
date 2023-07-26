package vec3

import (
	"math"
	"testing"
)

func TestVec3ScalarAdd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	y := Vec3ScalarAdd(x1, x2)
	true_y := Vec3{4.0, 5.0, 6.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} + 3.0 = {%f, %f, %f}; want {4.0, 5.0}",
			y.x,
			y.y,
			y.z,
		)
	}
}

func TestVec3ScalarSubtract(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	y := Vec3ScalarSubtract(x1, x2)
	true_y := Vec3{-2.0, -1.0, 0.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 0.0} - 3.0 = {%f, %f %f}; want {-2.0, -1.0, 0.0}",
			y.x,
			y.y,
			y.z,
		)
	}
}

func TestVec3ScalarProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	y := Vec3ScalarProd(x1, x2)
	true_y := Vec3{3.0, 6.0, 9.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} * 3.0 = {%f, %f, %f}; want {3.0, 6.0, 9.0}",
			y.x,
			y.y,
			y.z,
		)
	}
}

func TestVec3Add(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	y := Vec3Add(x1, x2)
	true_y := Vec3{5.0, -4.0, 5.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} + {1.0, 2.0, 5.0} = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}

func TestVec3Subtract(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	y := Vec3Subtract(x1, x2)
	true_y := Vec3{-3.0, 8.0, 1.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} - {1.0, 2.0, 2.0} = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}

func TestVec3ElementWiseProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	y := Vec3ElementWiseProd(x1, x2)
	true_y := Vec3{4.0, -12.0, 6.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} .* {1.0, 2.0, 2.0} = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}

func TestVec3DotProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	y := Vec3DotProd(x1, x2)
	true_y := -2.0

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} . {1.0, 2.0, 2.0} = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec3FloorPos(t *testing.T) {
	x := Vec3{1.9, 2.3, 3.5}

	y := Vec3Floor(x)
	true_y := Vec3{1.0, 2.0, 3.0}

	if y != true_y {
		t.Errorf(
			"Floor({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}

func TestVec3FloorNeg(t *testing.T) {
	x := Vec3{-1.9, -2.3, -0.1}

	y := Vec3Floor(x)
	true_y := Vec3{-2.0, -3.0, -1.0}

	if y != true_y {
		t.Errorf(
			"Floor({-1.9, -2.3, -0.1}) = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}
func TestVec3FractPos(t *testing.T) {
	x := Vec3{1.9, 2.3, 3.5}

	y := Vec3Fract(x)
	true_y := Vec3{0.9, 0.3, 0.5}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) || math.Abs(y.z-true_y.z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
			true_y.x-y.x,
			true_y.y-y.y,
			true_y.z-y.z,
		)
	}
}

func TestVec3FractNeg(t *testing.T) {
	x := Vec3{-1.9, -2.3, -0.1}

	y := Vec3Fract(x)
	true_y := Vec3{0.1, 0.7, 0.9}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) || math.Abs(y.z-true_y.z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
			true_y.x-y.x,
			true_y.y-y.y,
			true_y.z-y.z,
		)
	}
}

func TestVec3LengthSquared(t *testing.T) {
	x := Vec3{3.0, 4.0, 2.0}

	y := Vec3LengthSquared(x)
	true_y := 29.0

	if y != true_y {
		t.Errorf(
			"len({3.0, 4.0, 2.0})^2 = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec3Length(t *testing.T) {
	x := Vec3{3.0, 4.0, 0.0}

	y := Vec3Length(x)
	true_y := 5.0

	if y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec3Abs(t *testing.T) {
	x := Vec3{3.0, -4.0, 0.0}

	y := Vec3Abs(x)
	true_y := Vec3{3.0, 4.0, 0.0}

	if y != true_y {
		t.Errorf(
			"Abs({-1.9, -2.3, -0.1}) = {%f, %f, %f}; want {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
		)
	}
}

func TestVec3Unit(t *testing.T) {
	x := Vec3{1.0, 2.0, 2.0}

	y := Vec3Unit(x)
	true_y := Vec3{1.0 / 3.0, 2.0 / 3.0, 2.0 / 3.0}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) || math.Abs(y.z-true_y.z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			y.x,
			y.y,
			y.z,
			true_y.x,
			true_y.y,
			true_y.z,
			true_y.x-y.x,
			true_y.y-y.y,
			true_y.z-y.z,
		)
	}
}
