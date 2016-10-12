package base

import (
	"math"
	"math/rand"
)

// Sphere ...
type Sphere struct {
	center Vec3
	radius float64
	mat    Material
}

// NewSphere constructs a new sphere object with the specified parameters.
func NewSphere(center Vec3, radius float64, mat Material) *Sphere {
	return &Sphere{center, radius, mat}
}

// RandomInUnitSphere returns a point in the unit sphere
func RandomInUnitSphere() Vec3 {
	v2 := NewVec3(1, 1, 1)
	for {
		p := NewVec3(rand.Float64(), rand.Float64(), rand.Float64()).
			MultiplyScalar(2.0).Subtract(v2)
		if p.SquaredMagnitude() < 1.0 {
			return p
		}
	}
}

// Hit returns the value t the ray intersects with a point on the sphere,
// otherwise returns -1.0.
func (s *Sphere) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin().Subtract(s.center)
	a := r.Direction().Dot(r.Direction())
	b := 2 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		if t := (-b - math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			rec.t = t
			rec.p = r.PointAt(t)
			rec.normal = rec.Point().Subtract(s.center).DivideScalar(s.radius)
			rec.mat = s.mat
			return true
		}
		if t := (-b + math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			rec.t = t
			rec.p = r.PointAt(t)
			rec.normal = rec.Point().Subtract(s.center).DivideScalar(s.radius)
			rec.mat = s.mat
			return true
		}
	}
	return false
}
