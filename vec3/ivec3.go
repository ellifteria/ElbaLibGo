package vec3

import "math"

type IVec3 struct {
	X int
	Y int
	Z int
}

func IVec3ScalarAdd(lhs IVec3, rhs int) IVec3 {
	return IVec3{
		lhs.X + rhs,
		lhs.Y + rhs,
		lhs.Z + rhs,
	}
}

func IVec3ScalarSubtract(lhs IVec3, rhs int) IVec3 {
	return IVec3ScalarAdd(lhs, -rhs)
}

func IVec3ScalarProd(lhs IVec3, rhs int) IVec3 {
	return IVec3{
		lhs.X * rhs,
		lhs.Y * rhs,
		lhs.Z * rhs,
	}
}

func IVec3ScalarPow(lhs IVec3, rhs int) IVec3 {
	return IVec3{
		int(math.Pow(float64(lhs.X), float64(rhs))),
		int(math.Pow(float64(lhs.Y), float64(rhs))),
		int(math.Pow(float64(lhs.Z), float64(rhs))),
	}
}

func IVec3Add(lhs IVec3, rhs IVec3) IVec3 {
	return IVec3{
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
		lhs.Z + rhs.Z,
	}
}

func IVec3Subtract(lhs IVec3, rhs IVec3) IVec3 {
	return IVec3Add(lhs, IVec3ScalarProd(rhs, -1))
}

func IVec3ElementWiseProd(lhs IVec3, rhs IVec3) IVec3 {
	return IVec3{
		lhs.X * rhs.X,
		lhs.Y * rhs.Y,
		lhs.Z * rhs.Z,
	}
}

func IVec3DotProd(lhs IVec3, rhs IVec3) int {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func IVec3Abs(v IVec3) IVec3 {
	return IVec3{
		int(math.Abs(float64(v.X))),
		int(math.Abs(float64(v.Y))),
		int(math.Abs(float64(v.Z))),
	}
}

func IVec3LengthSquared(v IVec3) float64 {
	return float64(IVec3DotProd(v, v))
}

func IVec3Length(v IVec3) float64 {
	return math.Sqrt(IVec3LengthSquared(v))
}

func IVec3ToVec3(v IVec3) Vec3 {
	return Vec3{
		float64(v.X),
		float64(v.Y),
		float64(v.Z),
	}
}
