package primitives

import (
	"fmt"
	"math"
)

// Unit vectors
var (
	UnitX = NewVec3(1.0, 0, 0)
	UnitY = NewVec3(0, 1.0, 0)
	UnitZ = NewVec3(0, 0, 1.0)
)

// Vec3 ...
type Vec3 struct {
	vec []float64
}

// Vec returns the slice that represents our vector. This is a convenient
// method if we need to iterate through all the points that make up our
// vector.
func (v Vec3) Vec() []float64 {
	return v.vec
}

// X returns the x-value of the vector.
func (v Vec3) X() float64 {
	return v.vec[0]
}

// Y returns the y-value of the vector.
func (v Vec3) Y() float64 {
	return v.vec[1]
}

// Z returns the z-value of the vector.
func (v Vec3) Z() float64 {
	return v.vec[2]
}

// NewVec3 returns a vector object with the specified coordinates.
func NewVec3(x, y, z float64) Vec3 {
	vec := []float64{x, y, z}
	return Vec3{vec}
}

// Add returns the sum of two vectors.
func (v Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v.X()+v2.X(), v.Y()+v2.Y(), v.Z()+v2.Z())
}

// Subtract returns the difference of two vectors.
func (v Vec3) Subtract(v2 Vec3) Vec3 {
	return NewVec3(v.X()-v2.X(), v.Y()-v2.Y(), v.Z()-v2.Z())
}

// Multiply returns the product of two vectors.
func (v Vec3) Multiply(v2 Vec3) Vec3 {
	return NewVec3(v.X()*v2.X(), v.Y()*v2.Y(), v.Z()*v2.Z())
}

// Divide returns the product of two vectors.
func (v Vec3) Divide(v2 Vec3) Vec3 {
	return NewVec3(v.X()/v2.X(), v.Y()/v2.Y(), v.Z()/v2.Z())
}

// Normalize normalizes the vector.
func (v Vec3) Normalize() Vec3 {
	magnitude := math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
	// v.X() = v.X() / magnitude
	// v.Y() = v.Y() / magnitude
	// v.Z() = v.Z() / magnitude
	return NewVec3(v.X()/magnitude, v.Y()/magnitude, v.Z()/magnitude)
}

// Dot returns the dot product between two vectors.
func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X()*v2.X() + v.Y()*v2.Y() + v.Z()*v2.Z()
}

// Cross returns the cross product between two vectors.
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return NewVec3(v.Y()*v2.Z()-v.Z()*v2.Y(), -(v.X()*v2.Z() - v.Z()*v2.X()), v.X()*v2.Y()-v.Y()*v2.X())
}

// Magnitude returns the magnitude of the vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
}

// SquaredMagnitude returns the squared magnitude of the vector.
func (v Vec3) SquaredMagnitude() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

// MultiplyScalar returns a vector multiplied by s.
func (v Vec3) MultiplyScalar(s float64) Vec3 {
	return NewVec3(v.X()*s, v.Y()*s, v.Z()*s)
}

// DivideScalar returns a vector divided by s.
func (v Vec3) DivideScalar(s float64) Vec3 {
	return NewVec3(v.X()/s, v.Y()/s, v.Z()/s)
}

// Reflect returns the vector in which the incidental vector reflects relative
// to the normal.
func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Subtract(n.MultiplyScalar(2 * v.Dot(n)))
}

// Refract returns the vector from refraction according to Snell's Law.
func (v Vec3) Refract(n Vec3, niOverNt float64) (bool, Vec3) {
	uv := v.Normalize()
	dt := uv.Dot(n)
	discriminant := 1.0 - (niOverNt * niOverNt * (1 - dt*dt))
	if discriminant > 0 {
		refracted := uv.Subtract(n.MultiplyScalar(dt)).MultiplyScalar(niOverNt).
			Subtract(n.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, Vec3{}
}

func (v Vec3) String() string {
	return fmt.Sprintf("{ x:%f, y:%f, z:%f }", v.X(), v.Y(), v.Z())
}
