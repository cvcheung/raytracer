package textures

import (
	"math"
	"raytracer/primitives"
)

// Checkered ...
type Checkered struct {
	even, odd Color
}

// NewCheckered ...
func NewCheckered(even, odd Color) Checkered {
	return Checkered{even, odd}
}

// GetColor returns one of two alternating colors.
func (c Checkered) GetColor(u, v float64, p primitives.Vec3) Color {
	sines := math.Sin(10*p.X()) * math.Sin(10*p.Y()) * math.Sin(10*p.Z())
	if sines < 0 {
		return c.odd.GetColor(u, v, p)
	}
	return c.even.GetColor(u, v, p)
}
