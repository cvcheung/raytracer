package objects

import (
	"raytracer/materials"
	"raytracer/primitives"
)

// Triangle ...
type Triangle struct {
	v1, v2, v3 primitives.Vec3
	mat        materials.Material
}

// Hit returns true if a ray intersects with the triangle and stores the result
// in the passed record.
func (s *Triangle) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	return false
}
