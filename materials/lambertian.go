package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/utils"
)

// Lambertian defines the matte material.
type Lambertian struct {
	albedo textures.Texture
}

// NewLambertian returns a new matte definition.
func NewLambertian(color textures.Texture) Lambertian {
	return Lambertian{color}
}

// Scatter randomly bounces the ray to give a fairly accurate representation of
// the diffuse effect.
func (l Lambertian) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord) (bool, *primitives.Ray) {
	attenuation.Update(l.albedo.GetColor(0, 0, rec.Point()))
	target := rec.Point().Add(rec.Normal()).Add(utils.RandomInUnitSphere())
	return true, primitives.NewRay(rec.Point(), target.Subtract(rec.Point()))
}
