package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
	"raytracer/utils"
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

// Hit returns true if a ray intersects with the sphere and stores the result in
// the passed record.
func (s *Sphere) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	oc := r.Origin().Subtract(s.center)
	a := r.Direction().Dot(r.Direction())
	b := 2 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		if t := (-b - math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			p := r.PointAt(t)
			normal := p.Subtract(s.center).DivideScalar(s.radius)
			u, v := utils.GetSphereUV(normal)
			rec.UpdateRecord(t, u, v, p, normal, s.mat)
			return true
		}
		if t := (-b + math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			p := r.PointAt(t)
			normal := p.Subtract(s.center).DivideScalar(s.radius)
			u, v := utils.GetSphereUV(normal)
			rec.UpdateRecord(t, u, v, p, normal, s.mat)
			return true
		}
	}
	return false
}

// BoundingBox returns the AABB for a sphere.
func (s *Sphere) BoundingBox(t0, t1 float64) (bool, *AABB) {
	radii := primitives.NewVec3(s.radius, s.radius, s.radius)
	return true, NewAABB(s.center.Subtract(radii), s.center.Add(radii))
}
