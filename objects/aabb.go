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

// NewEmptyAABB returns a new AABB used for temporary variables.
func NewEmptyAABB() *AABB {
	return &AABB{}
}

// CopyAABB update the current record with the fields another record.
func (a *AABB) CopyAABB(a2 *AABB) {
	a.min = a2.min
	a.max = a2.max
}

// Min returns the lower end that defines our AABB.
func (a *AABB) Min() primitives.Vec3 {
	return a.min
}

// Max returns the lower end that defines our AABB.
func (a *AABB) Max() primitives.Vec3 {
	return a.max
}

// Area returns the area contained by the bounding box
func (a *AABB) Area() float64 {
	x := a.max.X() - a.min.X()
	y := a.max.Y() - a.min.Y()
	z := a.max.Z() - a.min.Z()
	return 2 * (x*y + x*z + y*z)
}

// LongestAxis returns the longest axis of our AABB.
func (a *AABB) LongestAxis() float64 {
	max := math.Max(
		math.Max(a.max.X()-a.min.X(), a.max.Y()-a.min.Y()),
		a.max.Z()-a.min.Z())
	if max == a.max.X() {
		return 0
	} else if max == a.max.Y() {
		return 1
	}
	return 0.5
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
