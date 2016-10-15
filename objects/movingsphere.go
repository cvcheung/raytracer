package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
	"raytracer/utils"
)

// MovingSphere ...
type MovingSphere struct {
	c0, c1         primitives.Vec3
	t0, t1, radius float64
	mat            materials.Material
}

// NewMovingSphere constructs a new moving sphere object with the specified
// parameters.
func NewMovingSphere(c0, c1 primitives.Vec3, t0, t1, radius float64, mat materials.Material) *MovingSphere {
	return &MovingSphere{c0, c1, t0, t1, radius, mat}
}

func (s *MovingSphere) center(time float64) primitives.Vec3 {
	t := (time - s.t0) / (s.t1 - s.t0)
	return s.c0.Add(s.c1.Subtract(s.c0).MultiplyScalar(t))
}

// Hit returns true if a ray intersects with the sphere and stores the result in
// the passed record.
func (s *MovingSphere) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	oc := r.Origin().Subtract(s.center(r.Time()))
	a := r.Direction().Dot(r.Direction())
	b := 2 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		if t := (-b - math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			p := r.PointAt(t)
			normal := p.Subtract(s.center(r.Time())).DivideScalar(s.radius)
			u, v := utils.GetSphereUV(normal)
			rec.UpdateRecord(t, u, v, p, normal, s.mat)
			return true
		}
		if t := (-b + math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			p := r.PointAt(t)
			normal := p.Subtract(s.center(r.Time())).DivideScalar(s.radius)
			u, v := utils.GetSphereUV(normal)
			rec.UpdateRecord(t, u, v, p, normal, s.mat)
			return true
		}
	}
	return false
}

// BoundingBox returns the AABB for a moving sphere.
func (s *MovingSphere) BoundingBox(t0, t1 float64) (bool, *AABB) {
	radii := primitives.NewVec3(s.radius, s.radius, s.radius)
	c0 := s.center(t0)
	c1 := s.center(t1)
	box0 := NewAABB(c0.Subtract(radii), c0.Add(radii))
	box1 := NewAABB(c1.Subtract(radii), c1.Add(radii))
	box := SurroundingBox(box0, box1)
	return true, box
}
