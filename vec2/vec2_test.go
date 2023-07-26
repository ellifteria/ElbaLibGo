package vec2

import (
	"math"
	"testing"
)

func TestVec2ScalarAdd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	Y := Vec2ScalarAdd(x1, x2)
	true_y := Vec2{4.0, 5.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} + 3.0 = {%f, %f}; want {4.0, 5.0}",
			Y.X,
			Y.Y,
		)
	}
}

func TestVec2ScalarSubtract(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	Y := Vec2ScalarSubtract(x1, x2)
	true_y := Vec2{-2.0, -1.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} - 3.0 = {%f, %f}; want {-2.0, -1.0}",
			Y.X,
			Y.Y,
		)
	}
}

func TestVec2ScalarProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	Y := Vec2ScalarProd(x1, x2)
	true_y := Vec2{3.0, 6.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} * 3.0 = {%f, %f}; want {3.0, 6.0}",
			Y.X,
			Y.Y,
		)
	}
}

func TestVec2Add(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	Y := Vec2Add(x1, x2)
	true_y := Vec2{5.0, -4.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} + {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}

func TestVec2Subtract(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	Y := Vec2Subtract(x1, x2)
	true_y := Vec2{-3.0, 8.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} - {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}

func TestVec2ElementWiseProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	Y := Vec2ElementWiseProd(x1, x2)
	true_y := Vec2{4.0, -12.0}

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} .* {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}

func TestVec2DotProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	Y := Vec2DotProd(x1, x2)
	true_y := -8.0

	if Y != true_y {
		t.Errorf(
			"{1.0, 2.0} . {1.0, 2.0} = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec2FloorPos(t *testing.T) {
	X := Vec2{1.9, 2.3}

	Y := Vec2Floor(X)
	true_y := Vec2{1.0, 2.0}

	if Y != true_y {
		t.Errorf(
			"Floor({1.9, 2.3}) = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}

func TestVec2FloorNeg(t *testing.T) {
	X := Vec2{-1.9, -2.3}

	Y := Vec2Floor(X)
	true_y := Vec2{-2.0, -3.0}

	if Y != true_y {
		t.Errorf(
			"Floor({-1.9, -2.3}) = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}
func TestVec2FractPos(t *testing.T) {
	X := Vec2{1.9, 2.3}

	Y := Vec2Fract(X)
	true_y := Vec2{0.9, 0.3}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
		)
	}
}

func TestVec2FractNeg(t *testing.T) {
	X := Vec2{-1.9, -2.3}

	Y := Vec2Fract(X)
	true_y := Vec2{0.1, 0.7}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({-1.9, -2.3}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
		)
	}
}

func TestVec2LengthSquared(t *testing.T) {
	X := Vec2{3.0, 4.0}

	Y := Vec2LengthSquared(X)
	true_y := 25.0

	if Y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec2Length(t *testing.T) {
	X := Vec2{3.0, 4.0}

	Y := Vec2Length(X)
	true_y := 5.0

	if Y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			Y,
			true_y,
		)
	}
}

func TestVec2Abs(t *testing.T) {
	X := Vec2{3.0, -4.0}

	Y := Vec2Abs(X)
	true_y := Vec2{3.0, 4.0}

	if Y != true_y {
		t.Errorf(
			"Abs({-1.9, -2.3, -0.1}) = {%f, %f}; want {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
		)
	}
}

func TestVec2Unit(t *testing.T) {
	X := Vec2{3.0, 4.0}

	Y := Vec2Unit(X)
	true_y := Vec2{3.0 / 5.0, 4.0 / 5.0}

	if math.Abs(Y.X-true_y.X) > 1*math.Pow(10, -10) || math.Abs(Y.Y-true_y.Y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			Y.X,
			Y.Y,
			true_y.X,
			true_y.Y,
			true_y.X-Y.X,
			true_y.Y-Y.Y,
		)
	}
}
