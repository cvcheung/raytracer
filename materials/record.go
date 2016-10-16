package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
)

// HitRecord is a simple record to that records the information regarding where
// the ray hit.
type HitRecord struct {
	t, u, v   float64
	p, normal primitives.Vec3
	specular  textures.Color
	mat       Material
}

// NewRecord returns a new hit record with the following information.
func NewRecord(t, u, v float64, p, normal primitives.Vec3, mat Material) *HitRecord {
	return &HitRecord{t, u, v, p, normal, textures.Black, mat}
}

// UpdateRecord modifies a record with new fields.
func (rec *HitRecord) UpdateRecord(t, u, v float64, p, normal primitives.Vec3, mat Material) {
	rec.t = t
	rec.u = u
	rec.v = v
	rec.p = p
	rec.normal = normal
	rec.mat = mat
}

// SetSpecular ...
func (rec *HitRecord) SetSpecular(c textures.Color) {
	rec.specular = c
}

// Specular ...
func (rec *HitRecord) Specular() textures.Color {
	return rec.specular
}

// T returns the t value that caused the ray to intersect the object.
func (rec *HitRecord) T() float64 {
	return rec.t
}

// U returns the horizontal coordinate of the scene.
func (rec *HitRecord) U() float64 {
	return rec.u
}

// V returns the vertical coordinate of the scene.
func (rec *HitRecord) V() float64 {
	return rec.v
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
	rec.u = rec2.u
	rec.v = rec2.v
	rec.p = rec2.p
	rec.normal = rec2.normal
	rec.mat = rec2.mat
}
