package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
)

// AABB is an axis aligned bounding box used to accelerate the performance of
// our ray tracer.
type AABB struct {
	min, max primitives.Vec3
}

// NewAABB returns a new AABB specified by the the two vectors passed in.
func NewAABB(a, b primitives.Vec3) AABB {
	return AABB{a, b}
}

// Min returns the lower end that defines our AABB.
func (a AABB) Min() primitives.Vec3 {
	return a.min
}

// Max returns the lower end that defines our AABB.
func (a AABB) Max() primitives.Vec3 {
	return a.max
}

// Hit ...
func (a AABB) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	for i := 0; i < 3; i++ {
		t0 := math.Min((a.min.X()-r.Origin().X())/r.Direction().X(),
			(a.max.X()-r.Origin().X())/r.Direction().X())
		t1 := math.Max((a.min.X()-r.Origin().X())/r.Direction().X(),
			(a.max.X()-r.Origin().X())/r.Direction().X())
		tMin = math.Max(t0, tMin)
		tMax = math.Min(t1, tMax)
		if tMax <= tMin {
			return false
		}
	}
	return true
}
