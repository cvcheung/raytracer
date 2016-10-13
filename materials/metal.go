package materials

import (
	"math/rand"
	"raytracer/primitives"
	"raytracer/utils"
)

// Metal defines the metallic material. Fuzz is utilized to determine how sharp
// rays are reflected.
type Metal struct {
	albedo primitives.Color
	fuzz   float64
}

// NewMetal returns a new metal definition. Fuzz is bounded to less than or
// equal to 1.
func NewMetal(color primitives.Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{color, fuzz}
}

// NewRandomMetal returns a random metal definition using Go's builtin PRNG.
func NewRandomMetal() Metal {
	color := primitives.NewColor(0.5*(1+rand.Float64()),
		0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()))
	fuzz := 0.5 * rand.Float64()
	return Metal{color, fuzz}
}

// Color returns the base colors of the metal definition.
func (m Metal) Color() primitives.Color {
	return m.albedo
}

// Scatter calculates the incidental reflected ray if there is a reflection.
func (m Metal) Scatter(rayIn *primitives.Ray, rec *HitRecord) (bool, *primitives.Ray) {
	// reflected := rayIn.Direction().Normalize().Reflect(rec.Normal())
	reflected := utils.Reflect(rayIn.Direction().Normalize(), rec.Normal())
	scattered := primitives.NewRay(rec.Point(),
		reflected.Add(utils.RandomInUnitSphere().MultiplyScalar(m.fuzz)))
	return scattered.Direction().Dot(rec.normal) > 0, scattered
}
