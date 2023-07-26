package vec2

import "math"

type Vec2 struct {
	x float64
	y float64
}

func Vec2ScalarAdd(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		lhs.x + rhs,
		lhs.y + rhs,
	}
}

func Vec2ScalarSubtract(lhs Vec2, rhs float64) Vec2 {
	return Vec2ScalarAdd(lhs, -rhs)
}

func Vec2ScalarProd(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		lhs.x * rhs,
		lhs.y * rhs,
	}
}

func Vec2ScalarPow(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		math.Pow(lhs.x, rhs),
		math.Pow(lhs.y, rhs),
	}
}

func Vec2Add(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2{
		lhs.x + rhs.x,
		lhs.y + rhs.y,
	}
}

func Vec2Subtract(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2Add(lhs, Vec2ScalarProd(rhs, -1))
}

func Vec2ElementWiseProd(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2{
		lhs.x * rhs.x,
		lhs.y * rhs.y,
	}
}

func Vec2DotProd(lhs Vec2, rhs Vec2) float64 {
	return lhs.x*rhs.x + lhs.y*rhs.y
}

func Vec2Abs(v Vec2) Vec2 {
	return Vec2{
		math.Abs(v.x),
		math.Abs(v.y),
	}
}

func Vec2Unit(v Vec2) Vec2 {
	return Vec2ScalarProd(v, 1/Vec2Length(v))
}

func Vec2Floor(v Vec2) Vec2 {
	return Vec2{
		math.Floor(v.x),
		math.Floor(v.y),
	}
}

func Vec2Fract(v Vec2) Vec2 {
	return Vec2Subtract(v, Vec2Floor(v))
}

func Vec2LengthSquared(v Vec2) float64 {
	return Vec2DotProd(v, v)
}

func Vec2Length(v Vec2) float64 {
	return math.Sqrt(Vec2LengthSquared(v))
}
