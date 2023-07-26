package vec3

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

func Vec3ScalarAdd(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		lhs.x + rhs,
		lhs.y + rhs,
		lhs.z + rhs,
	}
}

func Vec3ScalarSubtract(lhs Vec3, rhs float64) Vec3 {
	return Vec3ScalarAdd(lhs, -rhs)
}

func Vec3ScalarProd(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		lhs.x * rhs,
		lhs.y * rhs,
		lhs.z * rhs,
	}
}

func Vec3ScalarPow(lhs Vec3, rhs float64) Vec3 {
	return Vec3{
		math.Pow(lhs.x, rhs),
		math.Pow(lhs.y, rhs),
		math.Pow(lhs.z, rhs),
	}
}

func Vec3Add(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{
		lhs.x + rhs.x,
		lhs.y + rhs.y,
		lhs.z + rhs.z,
	}
}

func Vec3Subtract(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3Add(lhs, Vec3ScalarProd(rhs, -1))
}

func Vec3ElementWiseProd(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{
		lhs.x * rhs.x,
		lhs.y * rhs.y,
		lhs.z * rhs.z,
	}
}

func Vec3DotProd(lhs Vec3, rhs Vec3) float64 {
	return lhs.x*rhs.x + lhs.y*rhs.y + lhs.z*rhs.z
}

func Vec3Abs(v Vec3) Vec3 {
	return Vec3{
		math.Abs(v.x),
		math.Abs(v.y),
		math.Abs(v.z),
	}
}

func Vec3Unit(v Vec3) Vec3 {
	return Vec3ScalarProd(v, 1/Vec3Length(v))
}

func Vec3Floor(v Vec3) Vec3 {
	return Vec3{
		math.Floor(v.x),
		math.Floor(v.y),
		math.Floor(v.z),
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
