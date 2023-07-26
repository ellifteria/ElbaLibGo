package vec3

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func Vec3ScalarAdd(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		lhs.X + rhs,
		lhs.Y + rhs,
		lhs.Z + rhs,
	}
}

func Vec3ScalarSubtract(lhs Vec3, rhs float64) Vec3 {
	return Vec3ScalarAdd(lhs, -rhs)
}

func Vec3ScalarProd(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		lhs.X * rhs,
		lhs.Y * rhs,
		lhs.Z * rhs,
	}
}

func Vec3ScalarPow(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		math.Pow(lhs.X, rhs),
		math.Pow(lhs.Y, rhs),
		math.Pow(lhs.Z, rhs),
	}
}

func Vec3Add(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
		lhs.Z + rhs.Z,
	}
}

func Vec3Subtract(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3Add(lhs, Vec3ScalarProd(rhs, -1))
}

func Vec3ElementWiseProd(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{
		lhs.X * rhs.X,
		lhs.Y * rhs.Y,
		lhs.Z * rhs.Z,
	}
}

func Vec3DotProd(lhs Vec3, rhs Vec3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func Vec3Abs(v Vec3) Vec3 {
	return Vec3{
		math.Abs(v.X),
		math.Abs(v.Y),
		math.Abs(v.Z),
	}
}

func Vec3Unit(v Vec3) Vec3 {
	return Vec3ScalarProd(v, 1/Vec3Length(v))
}

func Vec3Floor(v Vec3) Vec3 {
	return Vec3{
		math.Floor(v.X),
		math.Floor(v.Y),
		math.Floor(v.Z),
	}
}

func Vec3Fract(v Vec3) Vec3 {
	return Vec3Subtract(v, Vec3Floor(v))
}

func Vec3LengthSquared(v Vec3) float64 {
	return Vec3DotProd(v, v)
}

func Vec3Length(v Vec3) float64 {
	return math.Sqrt(Vec3LengthSquared(v))
}

func Vec3ToIVec3(v Vec3) IVec3 {
	return IVec3{
		int(v.X),
		int(v.Y),
		int(v.Z),
	}
}
