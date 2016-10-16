package materials

import (
	"math/rand"
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/utils"
)

// Metal defines the metallic material. Fuzz is utilized to determine how sharp
// rays are reflected.
type Metal struct {
	albedo textures.Texture
	fuzz   float64
}

// NewMetal returns a new metal definition. Fuzz is bounded to less than or
// equal to 1.
func NewMetal(color textures.Texture, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{color, fuzz}
}

// NewRandomMetal returns a random metal definition using Go's builtin PRNG.
func NewRandomMetal() Metal {
	color := textures.NewColor(0.5*(1+rand.Float64()),
		0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()))
	fuzz := 0.5 * rand.Float64()
	return Metal{color, fuzz}
}

// Scatter calculates the incidental reflected ray if there is a reflection.
func (m Metal) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, lights []Light) (bool, *primitives.Ray) {
	attenuation.Update(m.albedo.GetColor(0, 0, rec.Point()))
	reflected := rayIn.Direction().Normalize().Reflect(rec.Normal())
	scattered := primitives.NewRay(rec.Point(),
		reflected.Add(utils.RandomInUnitSphere().MultiplyScalar(m.fuzz)))
	return scattered.Direction().Dot(rec.normal) > 0, scattered
}

// Emitted is defined to implement the material interface.
func (m Metal) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}
