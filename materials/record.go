package materials

import (
	"raytracer/primitives"
)

// HitRecord is a simple record to that records the information regarding where
// the ray hit.
type HitRecord struct {
	t         float64
	p, normal primitives.Vec3
	mat       Material
}

// NewRecord returns a new hit record with the following information.
func NewRecord(t float64, p, normal primitives.Vec3, mat Material) *HitRecord {
	return &HitRecord{t, p, normal, mat}
}

// UpdateRecord modifies a record with new fields.
func (rec *HitRecord) UpdateRecord(t float64, p, normal primitives.Vec3, mat Material) {
	rec.t = t
	rec.p = p
	rec.normal = normal
	rec.mat = mat
}

// T returns the t value that caused the ray to intersect the object.
func (rec *HitRecord) T() float64 {
	return rec.t
}

// Point returns the the point at which the ray intersected the object.
func (rec *HitRecord) Point() primitives.Vec3 {
	return rec.p
}

// Normal returns the normal at which the ray intersected the object.
func (rec *HitRecord) Normal() primitives.Vec3 {
	return rec.normal
}

// Material returns the pointer to the material struct that defines the type of
// material that the ray hit.
func (rec *HitRecord) Material() Material {
	return rec.mat
}

// CopyRecord update the current record with the fields another record.
func (rec *HitRecord) CopyRecord(rec2 *HitRecord) {
	rec.t = rec2.t
	rec.p = rec2.p
	rec.normal = rec2.normal
	rec.mat = rec2.mat
}
