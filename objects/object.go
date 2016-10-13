package objects

import (
	"raytracer/materials"
	"raytracer/primitives"
)

// Object defines base class for 2/3D objects.
type Object interface {
	Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool
}
