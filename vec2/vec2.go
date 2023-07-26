package vec2

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func Vec2ScalarAdd(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		lhs.X + rhs,
		lhs.Y + rhs,
	}
}

func Vec2ScalarSubtract(lhs Vec2, rhs float64) Vec2 {
	return Vec2ScalarAdd(lhs, -rhs)
}

func Vec2ScalarProd(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		lhs.X * rhs,
		lhs.Y * rhs,
	}
}

func Vec2ScalarPow(lhs Vec2, rhs float64) Vec2 {
	return Vec2{
		math.Pow(lhs.X, rhs),
		math.Pow(lhs.Y, rhs),
	}
}

func Vec2Add(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2{
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
	}
}

func Vec2Subtract(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2Add(lhs, Vec2ScalarProd(rhs, -1))
}

func Vec2ElementWiseProd(lhs Vec2, rhs Vec2) Vec2 {
	return Vec2{
		lhs.X * rhs.X,
		lhs.Y * rhs.Y,
	}
}

func Vec2DotProd(lhs Vec2, rhs Vec2) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y
}

func Vec2Abs(v Vec2) Vec2 {
	return Vec2{
		math.Abs(v.X),
		math.Abs(v.Y),
	}
}

func Vec2Unit(v Vec2) Vec2 {
	return Vec2ScalarProd(v, 1/Vec2Length(v))
}

func Vec2Floor(v Vec2) Vec2 {
	return Vec2{
		math.Floor(v.X),
		math.Floor(v.Y),
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

func Vec2ToIVec2(v Vec2) IVec2 {
	return IVec2{
		int(v.X),
		int(v.Y),
	}
}
