package utils

import (
	"math"
	"math/rand"
	"raytracer/primitives"
)

// GetSphereUV corresponding u, v coordinates in the image plane.
func GetSphereUV(p primitives.Vec3) (float64, float64) {
	phi := math.Atan2(p.Z(), p.X())
	theta := math.Asin(p.Y())
	u := 1 - (phi+math.Pi)/(2*math.Pi)
	v := (theta + math.Pi/2) / math.Pi
	return u, v
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
