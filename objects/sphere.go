package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
)

// Sphere ...
type Sphere struct {
	center primitives.Vec3
	radius float64
	mat    materials.Material
}

// NewSphere constructs a new sphere object with the specified parameters.
func NewSphere(center primitives.Vec3, radius float64, mat materials.Material) *Sphere {
	return &Sphere{center, radius, mat}
}

// Hit returns the value t the ray intersects with a point on the sphere,
// otherwise returns -1.0.
func (s *Sphere) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	oc := r.Origin().Subtract(s.center)
	a := r.Direction().Dot(r.Direction())
	b := 2 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		if t := (-b - math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			normal := rec.Point().Subtract(s.center).DivideScalar(s.radius)
			rec.UpdateRecord(t, r.PointAt(t), normal, s.mat)
			return true
		}
		if t := (-b + math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			normal := rec.Point().Subtract(s.center).DivideScalar(s.radius)
			rec.UpdateRecord(t, r.PointAt(t), normal, s.mat)
			return true
		}
	}
	return false
}