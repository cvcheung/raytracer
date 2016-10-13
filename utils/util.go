package utils

import (
	"math"
	"math/rand"
	"raytracer/primitives"
)

// Reflect returns the vector in which the incidental vector reflects relative
// to the normal.
func Reflect(v, n primitives.Vec3) primitives.Vec3 {
	return v.Subtract(n.MultiplyScalar(2 * v.Dot(n)))
}

// Refract returns the vector from refraction according to Snell's Law.
func Refract(v, n primitives.Vec3, niOverNt float64) (bool, primitives.Vec3) {
	uv := v.Normalize()
	dt := uv.Dot(n)
	discriminant := 1.0 - (niOverNt * niOverNt * (1 - dt*dt))
	if discriminant > 0 {
		refracted := uv.Subtract(n.MultiplyScalar(dt)).MultiplyScalar(niOverNt).
			Subtract(n.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, primitives.Vec3{}
}

// RandomInUnitSphere returns a point in the unit sphere.
func RandomInUnitSphere() primitives.Vec3 {
	v2 := primitives.NewVec3(1, 1, 1)
	for {
		p := primitives.NewVec3(rand.Float64(), rand.Float64(), rand.Float64()).
			MultiplyScalar(2.0).Subtract(v2)
		if p.SquaredMagnitude() < 1.0 {
			return p
		}
	}
}

// RandomInUnitDisk returns a point in the unit disk.
func RandomInUnitDisk() primitives.Vec3 {
	v2 := primitives.NewVec3(1, 1, 0)
	for {
		p := primitives.NewVec3(rand.Float64(), rand.Float64(), 0).
			MultiplyScalar(2.0).Subtract(v2)
		if p.SquaredMagnitude() < 1.0 {
			return p
		}
	}
}

// Schlick angle reflection approximation for s a formula for approximating the
// contribution of the Fresnel factor in the specular reflection of light from
// a non-conducting interface (surface) between two media.
// [https://en.wikipedia.org/wiki/Schlick%27s_approximation]
func Schlick(cosine, reflectIdx float64) float64 {
	r0 := (1 - reflectIdx) / (1 + reflectIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
