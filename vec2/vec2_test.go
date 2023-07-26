package vec2

import (
	"math"
	"testing"
)

func TestVec2ScalarAdd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	y := Vec2ScalarAdd(x1, x2)
	true_y := Vec2{4.0, 5.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} + 3.0 = {%f, %f}; want {4.0, 5.0}",
			y.x,
			y.y,
		)
	}
}

func TestVec2ScalarSubtract(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	y := Vec2ScalarSubtract(x1, x2)
	true_y := Vec2{-2.0, -1.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} - 3.0 = {%f, %f}; want {-2.0, -1.0}",
			y.x,
			y.y,
		)
	}
}

func TestVec2ScalarProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := 3.0

	y := Vec2ScalarProd(x1, x2)
	true_y := Vec2{3.0, 6.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} * 3.0 = {%f, %f}; want {3.0, 6.0}",
			y.x,
			y.y,
		)
	}
}

func TestVec2Add(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	y := Vec2Add(x1, x2)
	true_y := Vec2{5.0, -4.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} + {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}

func TestVec2Subtract(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	y := Vec2Subtract(x1, x2)
	true_y := Vec2{-3.0, 8.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} - {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}

func TestVec2ElementWiseProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	y := Vec2ElementWiseProd(x1, x2)
	true_y := Vec2{4.0, -12.0}

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} .* {1.0, 2.0} = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}

func TestVec2DotProd(t *testing.T) {
	x1 := Vec2{1.0, 2.0}
	x2 := Vec2{4.0, -6.0}

	y := Vec2DotProd(x1, x2)
	true_y := -8.0

	if y != true_y {
		t.Errorf(
			"{1.0, 2.0} . {1.0, 2.0} = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec2FloorPos(t *testing.T) {
	x := Vec2{1.9, 2.3}

	y := Vec2Floor(x)
	true_y := Vec2{1.0, 2.0}

	if y != true_y {
		t.Errorf(
			"Floor({1.9, 2.3}) = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}

func TestVec2FloorNeg(t *testing.T) {
	x := Vec2{-1.9, -2.3}

	y := Vec2Floor(x)
	true_y := Vec2{-2.0, -3.0}

	if y != true_y {
		t.Errorf(
			"Floor({-1.9, -2.3}) = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}
func TestVec2FractPos(t *testing.T) {
	x := Vec2{1.9, 2.3}

	y := Vec2Fract(x)
	true_y := Vec2{0.9, 0.3}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
			true_y.x-y.x,
			true_y.y-y.y,
		)
	}
}

func TestVec2FractNeg(t *testing.T) {
	x := Vec2{-1.9, -2.3}

	y := Vec2Fract(x)
	true_y := Vec2{0.1, 0.7}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({-1.9, -2.3}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
			true_y.x-y.x,
			true_y.y-y.y,
		)
	}
}

func TestVec2LengthSquared(t *testing.T) {
	x := Vec2{3.0, 4.0}

	y := Vec2LengthSquared(x)
	true_y := 25.0

	if y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec2Length(t *testing.T) {
	x := Vec2{3.0, 4.0}

	y := Vec2Length(x)
	true_y := 5.0

	if y != true_y {
		t.Errorf(
			"len({3.0, 4.0})^2 = %f; want %f",
			y,
			true_y,
		)
	}
}

func TestVec2Abs(t *testing.T) {
	x := Vec2{3.0, -4.0}

	y := Vec2Abs(x)
	true_y := Vec2{3.0, 4.0}

	if y != true_y {
		t.Errorf(
			"Abs({-1.9, -2.3, -0.1}) = {%f, %f}; want {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
		)
	}
}

func TestVec2Unit(t *testing.T) {
	x := Vec2{3.0, 4.0}

	y := Vec2Unit(x)
	true_y := Vec2{3.0 / 5.0, 4.0 / 5.0}

	if math.Abs(y.x-true_y.x) > 1*math.Pow(10, -10) || math.Abs(y.y-true_y.y) > 1*math.Pow(10, -10) {
		t.Errorf(
			"Fract({1.9, 2.3, 3.5}) = {%f, %f}; want {%f, %f}; diff = {%f, %f}",
			y.x,
			y.y,
			true_y.x,
			true_y.y,
			true_y.x-y.x,
			true_y.y-y.y,
		)
	}
}
