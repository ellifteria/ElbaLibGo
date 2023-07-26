package vec3

import (
	"math"
	"testing"
)

func TestVec3ScalarAdd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	Y := Vec3ScalarAdd(x1, x2)
	true_y := Vec3{4.0, 5.0, 6.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} + 3.0 = {%f, %f, %f}; want {4.0, 5.0}",
			Y.X,
			Y.Y,
			Y.Z,
		)
	}
}

func TestVec3ScalarSubtract(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	Y := Vec3ScalarSubtract(x1, x2)
	true_y := Vec3{-2.0, -1.0, 0.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 0.0} - 3.0 = {%f, %f %f}; want {-2.0, -1.0, 0.0}",
			Y.X,
			Y.Y,
			Y.Z,
		)
	}
}

func TestVec3ScalarProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := 3.0

	Y := Vec3ScalarProd(x1, x2)
	true_y := Vec3{3.0, 6.0, 9.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} * 3.0 = {%f, %f, %f}; want {3.0, 6.0, 9.0}",
			Y.X,
			Y.Y,
			Y.Z,
		)
	}
}

func TestVec3Add(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	Y := Vec3Add(x1, x2)
	true_y := Vec3{5.0, -4.0, 5.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} + {1.0, 2.0, 5.0} = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}

func TestVec3Subtract(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	Y := Vec3Subtract(x1, x2)
	true_y := Vec3{-3.0, 8.0, 1.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} - {1.0, 2.0, 2.0} = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}

func TestVec3ElementWiseProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	Y := Vec3ElementWiseProd(x1, x2)
	true_y := Vec3{4.0, -12.0, 6.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} .* {1.0, 2.0, 2.0} = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}

func TestVec3DotProd(t *testing.T) {
	x1 := Vec3{1.0, 2.0, 3.0}
	x2 := Vec3{4.0, -6.0, 2.0}

	Y := Vec3DotProd(x1, x2)
	true_y := -2.0

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0, 3.0} . {1.0, 2.0, 2.0} = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec3FloorPos(t *testing.T) {
	X := Vec3{1.9, 2.3, 3.5}

	Y := Vec3Floor(X)
	true_y := Vec3{1.0, 2.0, 3.0}

	if Y != true_y {
		t.Errorf(
			"Floor({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}

func TestVec3FloorNeg(t *testing.T) {
	X := Vec3{-1.9, -2.3, -0.1}

	Y := Vec3Floor(X)
	true_y := Vec3{-2.0, -3.0, -1.0}

	if Y != true_y {
		t.Errorf(
			"Floor({-1.9, -2.3, -0.1}) = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}
func TestVec3FractPos(t *testing.T) {
	X := Vec3{1.9, 2.3, 3.5}

	Y := Vec3Fract(X)
	true_y := Vec3{0.9, 0.3, 0.5}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) || math.Abs(Y.Z-true_y.Z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
			true_y.Z-Y.Z,
		)
	}
}

func TestVec3FractNeg(t *testing.T) {
	X := Vec3{-1.9, -2.3, -0.1}

	Y := Vec3Fract(X)
	true_y := Vec3{0.1, 0.7, 0.9}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) || math.Abs(Y.Z-true_y.Z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
			true_y.Z-Y.Z,
		)
	}
}

func TestVec3LengthSquared(t *testing.T) {
	X := Vec3{3.0, 4.0, 2.0}

	Y := Vec3LengthSquared(X)
	true_y := 29.0

	if Y != true_y {
		t.Errorf(
			"len({3.0, 4.0, 2.0})^2 = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec3Length(t *testing.T) {
	X := Vec3{3.0, 4.0, 0.0}

	Y := Vec3Length(X)
	true_y := 5.0

	if Y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec3Abs(t *testing.T) {
	X := Vec3{3.0, -4.0, 0.0}

	Y := Vec3Abs(X)
	true_y := Vec3{3.0, 4.0, 0.0}

	if Y != true_y {
		t.Errorf(
			"Abs({-1.9, -2.3, -0.1}) = {%f, %f, %f}; want {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
		)
	}
}

func TestVec3Unit(t *testing.T) {
	X := Vec3{1.0, 2.0, 2.0}

	Y := Vec3Unit(X)
	true_y := Vec3{1.0 / 3.0, 2.0 / 3.0, 2.0 / 3.0}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) || math.Abs(Y.Z-true_y.Z) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f, %f}; want {%f, %f, %f}; diff = {%f, %f, %f}",
			Y.X,
			Y.Y,
			Y.Z,
			true_y.X,
			true_y.Y,
			true_y.Z,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
			true_y.Z-Y.Z,
		)
	}
}
