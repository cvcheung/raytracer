package primitives

import (
	"fmt"
	"math"
)

// Unit vectors
var (
	UnitX = Vec3{1.0, 0, 0}
	UnitY = Vec3{0, 1.0, 0}
	UnitZ = Vec3{0, 0, 1.0}
)

// Vec3 ...
type Vec3 struct {
	x, y, z float64
}

// X returns the x-value of the vector.
func (v Vec3) X() float64 {
	return v.x
}

// Y returns the y-value of the vector.
func (v Vec3) Y() float64 {
	return v.y
}

// Z returns the z-value of the vector.
func (v Vec3) Z() float64 {
	return v.z
}

// NewVec3 returns a vector object with the specified coordinates.
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Add returns the sum of two vectors.
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.x + v2.x, v.y + v2.y, v.z + v2.z}
}

// Subtract returns the difference of two vectors.
func (v Vec3) Subtract(v2 Vec3) Vec3 {
	return Vec3{v.x - v2.x, v.y - v2.y, v.z - v2.z}
}

// Multiply returns the product of two vectors.
func (v Vec3) Multiply(v2 Vec3) Vec3 {
	return Vec3{v.x * v2.x, v.y * v2.y, v.z * v2.z}
}

// Divide returns the product of two vectors.
func (v Vec3) Divide(v2 Vec3) Vec3 {
	return Vec3{v.x / v2.x, v.y / v2.y, v.z / v2.z}
}

// Normalize normalizes the vector.
func (v Vec3) Normalize() Vec3 {
	magnitude := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	// v.x = v.x / magnitude
	// v.y = v.y / magnitude
	// v.z = v.z / magnitude
	return Vec3{v.x / magnitude, v.y / magnitude, v.z / magnitude}
}

// Dot returns the dot product between two vectors.
func (v Vec3) Dot(v2 Vec3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

// Cross returns the cross product between two vectors.
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v.y*v2.z - v.z*v2.y, -(v.x*v2.z - v.z*v2.x), v.x*v2.y - v.y*v2.x}
}

// Magnitude returns the magnitude of the vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// SquaredMagnitude returns the squared magnitude of the vector.
func (v Vec3) SquaredMagnitude() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// MultiplyScalar returns a vector multiplied by s.
func (v Vec3) MultiplyScalar(s float64) Vec3 {
	return Vec3{v.x * s, v.y * s, v.z * s}
}

// DivideScalar returns a vector divided by s.
func (v Vec3) DivideScalar(s float64) Vec3 {
	return Vec3{v.x / s, v.y / s, v.z / s}
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
	return fmt.Sprintf("{ x:%f, y:%f, z:%f }", v.x, v.y, v.z)
}
