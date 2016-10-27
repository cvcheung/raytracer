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
func (l Lambertian) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, light Light, shadow bool) (bool, *primitives.Ray) {
	attenuation.Update(l.albedo.GetColor(0, 0, rec.Point()))
	target := rec.Point().Add(rec.Normal()).Add(utils.RandomInUnitSphere())
	return true, primitives.NewRay(rec.Point(), target.Subtract(rec.Point()))
}

// Emitted is defined to implement the material interface.
func (l Lambertian) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}

// GetAmbient ...
func (l Lambertian) GetAmbient() textures.Color {
	return textures.Black
}
