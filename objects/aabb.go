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
func NewAABB(a, b primitives.Vec3) *AABB {
	return &AABB{a, b}
}

// Min returns the lower end that defines our AABB.
func (a *AABB) Min() primitives.Vec3 {
	return a.min
}

// Max returns the lower end that defines our AABB.
func (a *AABB) Max() primitives.Vec3 {
	return a.max
}

// Hit ...
func (a *AABB) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	aMinVec := a.min.Vec()
	aMaxVec := a.max.Vec()
	originVec := r.Origin().Vec()
	directionVec := r.Direction().Vec()
	for i := 0; i < 3; i++ {
		// Proposed apporach
		// t0 := math.Min((aMinVec[i]-originVec[i])/directionVec[i],
		// 	(aMaxVec[i]-originVec[i])/directionVec[i])
		// t1 := math.Max((aMinVec[i]-originVec[i])/directionVec[i],
		// 	(aMaxVec[i]-originVec[i])/directionVec[i])
		// tMin = math.Max(t0, tMin)
		// tMax = math.Min(t1, tMax)
		// if tMax <= tMin {
		// 	return false
		// Alternative - Pixar approach
		invD := 1.0 / directionVec[i]
		t0 := (aMinVec[i] - originVec[i]) * invD
		t1 := (aMaxVec[i] - originVec[i]) * invD
		if invD < 0 {
			t0, t1 = t1, t0
		}
		if t0 > tMin {
			tMin = t0
		}
		if t1 < tMax {
			tMax = t1
		}
		if tMax < tMin {
			return false
		}
	}
	return true
}

// SurroundingBox calculates the encompassing AABB given two AABB.
func SurroundingBox(box0, box1 *AABB) *AABB {
	small := primitives.NewVec3(
		math.Min(box0.Min().X(), box1.Min().X()),
		math.Min(box0.Min().Y(), box1.Min().Y()),
		math.Min(box0.Min().Z(), box1.Min().Z()))
	big := primitives.NewVec3(
		math.Max(box0.Max().X(), box1.Max().X()),
		math.Max(box0.Max().Y(), box1.Max().Y()),
		math.Max(box0.Max().Z(), box1.Max().Z()))
	return NewAABB(small, big)
}

// BoxCompare ...
func BoxCompare(a, b Object) bool {
	return false
}
