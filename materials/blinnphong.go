package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/utils"
)

// Blinnphong defines a generic material that does the Blinn-Phong Shading.
type Blinnphong struct {
	ambient  textures.Texture
	diffuse  textures.Texture
	specular textures.Texture
	phong    float64
}

// NewBlinnphong returns a new material definition. Fuzz is bounded to less than or
// equal to 1.
func NewBlinnphong(ambient, diffuse, specular textures.Texture, phong float64) Blinnphong {
	return Blinnphong{ambient, specular, diffuse, phong}
}

// Scatter calculates the incidental reflected ray if there is a reflection.
func (b Blinnphong) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord) (bool, *primitives.Ray) {
	// attenuation.Update(b.albedo.GetColor(0, 0, rec.Point()))
	target := rec.Point().Add(rec.Normal()).Add(utils.RandomInUnitSphere())
	return true, primitives.NewRay(rec.Point(), target.Subtract(rec.Point()))
}

// Emitted is defined to implement the material interface.
func (b Blinnphong) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}
