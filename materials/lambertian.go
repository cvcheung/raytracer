package materials

import (
	"raytracer/primitives"
	"raytracer/utils"
)

// Lambertian defines the matte material.
type Lambertian struct {
	albedo primitives.Color
}

// NewLambertian returns a new matte definition.
func NewLambertian(color primitives.Color) Lambertian {
	return Lambertian{color}
}

// Color returns the base colors of the matte definition.
func (l Lambertian) Color() primitives.Color {
	return l.albedo
}

// Scatter randomly bounces the ray to give a fairly accurate representation of
// the diffuse effect.
func (l Lambertian) Scatter(rayIn *primitives.Ray, rec *HitRecord) (bool, *primitives.Ray) {
	target := rec.Point().Add(rec.Normal()).Add(utils.RandomInUnitSphere())
	return true, primitives.NewRay(rec.Point(), target.Subtract(rec.Point()))
}
