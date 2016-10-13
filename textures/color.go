package textures

import (
	"math/rand"
	"raytracer/primitives"
)

// Color is our color struct to support color adding. RGB values range from 0
// to 1.
type Color struct {
	R, G, B float64
}

// Generic exported colors
var (
	Red   = Color{1.0, 0, 0}
	White = Color{1.0, 1.0, 1.0}
	Blue  = Color{0.5, 0.7, 1.0}
	Black = Color{0.0, 0.0, 0.0}
)

// NewColor returns a color object with the specified values.
func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

// NewEmptyColor returns a color object with RGB all initialized to 0.
func NewEmptyColor() Color {
	return Color{0, 0, 0}
}

// NewRandomColor returns a random color.
func NewRandomColor() Color {
	return Color{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(),
		rand.Float64() * rand.Float64()}
}

// GetColor is used to have colors implment the texture interface. This allows
// solid colors to be textures.
func (c Color) GetColor(u, v float64, p primitives.Vec3) Color {
	return c
}

// Update modifies the color with new parameters.
func (c *Color) Update(c2 Color) {
	c.R = c2.R
	c.G = c2.G
	c.B = c2.B
}

// RGBA is taken from the go source code to implement the color.Color interface.
func (c Color) RGBA() (r, g, b, a uint32) {
	a = 0xffff
	r = uint32(c.R) * a
	g = uint32(c.G) * a
	b = uint32(c.B) * a
	return
}

// Add updates the color object with the sum of two colors.
func (c Color) Add(c2 Color) Color {
	return Color{c.R + c2.R, c.G + c2.G, c.B + c2.B}
}

// Multiply updates the color object the product of corresponding channels.
func (c Color) Multiply(c2 Color) Color {
	return Color{c.R * c2.R, c.G * c2.G, c.B * c2.B}
}

// AddScalar adds f to the all the color channels.
func (c Color) AddScalar(f float64) Color {
	return Color{c.R + f, c.G + f, c.B + f}
}

// MultiplyScalar multiplies f to the all the color channels.
func (c Color) MultiplyScalar(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f}
}

// DivideScalar divides f to the all the color channels.
func (c Color) DivideScalar(f float64) Color {
	return Color{c.R / f, c.G / f, c.B / f}
}

// Gradient returns a gradient of blue + white.
func Gradient(t float64) Color {
	return White.MultiplyScalar(1.0 - t).Add(Blue.MultiplyScalar(t))
}
