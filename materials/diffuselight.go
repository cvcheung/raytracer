package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/utils"
)

// DiffuseLight is the material used to make an area light.
type DiffuseLight struct {
	emit textures.Texture
}

// NewDiffuseLight returns a new matte definition.
func NewDiffuseLight(emit textures.Texture) DiffuseLight {
	return DiffuseLight{emit}
}

// Scatter randomly bounces the ray to give a fairly accurate representation of
// the diffuse effect.
func (d DiffuseLight) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, lights []Light) (bool, *primitives.Ray) {
	// attenuation.Update(l.albedo.GetColor(0, 0, rec.Point()))
	target := rec.Point().Add(rec.Normal()).Add(utils.RandomInUnitSphere())
	return true, primitives.NewRay(rec.Point(), target.Subtract(rec.Point()))
}

// Emitted is defined to implement the material interface.
func (d DiffuseLight) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return d.emit.GetColor(u, v, p)
}
