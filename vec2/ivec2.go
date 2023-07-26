package vec2

import "math"

type IVec2 struct {
	X int
	Y int
}

func IVec2ScalarAdd(lhs IVec2, rhs int) IVec2 {
	return IVec2{
		lhs.X + rhs,
		lhs.Y + rhs,
	}
}

func IVec2ScalarSubtract(lhs IVec2, rhs int) IVec2 {
	return IVec2ScalarAdd(lhs, -rhs)
}

func IVec2ScalarProd(lhs IVec2, rhs int) IVec2 {
	return IVec2{
		lhs.X * rhs,
		lhs.Y * rhs,
	}
}

func IVec2ScalarPow(lhs IVec2, rhs int) IVec2 {
	return IVec2{
		int(math.Pow(float64(lhs.X), float64(rhs))),
		int(math.Pow(float64(lhs.Y), float64(rhs))),
	}
}

func IVec2Add(lhs IVec2, rhs IVec2) IVec2 {
	return IVec2{
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
	}
}

func IVec2Subtract(lhs IVec2, rhs IVec2) IVec2 {
	return IVec2Add(lhs, IVec2ScalarProd(rhs, -1))
}

func IVec2ElementWiseProd(lhs IVec2, rhs IVec2) IVec2 {
	return IVec2{
		lhs.X * rhs.X,
		lhs.Y * rhs.Y,
	}
}

func IVec2DotProd(lhs IVec2, rhs IVec2) int {
	return lhs.X*rhs.X + lhs.Y*rhs.Y
}

func IVec2Abs(v IVec2) IVec2 {
	return IVec2{
		int(math.Abs(float64(v.X))),
		int(math.Abs(float64(v.Y))),
	}
}

func IVec2LengthSquared(v IVec2) float64 {
	return float64(IVec2DotProd(v, v))
}

func IVec2Length(v IVec2) float64 {
	return math.Sqrt(IVec2LengthSquared(v))
}

func IVec2ToVec2(v IVec2) Vec2 {
	return Vec2{
		float64(v.X),
		float64(v.Y),
	}
}
