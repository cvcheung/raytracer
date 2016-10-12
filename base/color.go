package base

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
)

// NewColor returns a color object with the specified values.
func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

// RGBA is taken from the go source code to implement the color.Color interface.
func (c Color) RGBA() (r, g, b, a uint32) {
	a = 0xffff
	r = uint32(c.R) * a
	g = uint32(c.G) * a
	b = uint32(c.B) * a
	return
}

func max(v1, v2 uint8) uint8 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func min(v1, v2 uint8) uint8 {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Add updates the color object with the sum of two colors.
// TODO Decide to use pointers or not.
// The reasoning behind using a pointer is that our film is storing color
// entries. By utilizing a pointer we can directly modify the color value
// already there without need for reassignment.
func (c Color) Add(c2 Color) Color {
	// c.R = c.R + c2.R
	// c.G = c.G + c2.G
	// c.B = c.B + c2.B
	return Color{c.R + c2.R, c.G + c2.G, c.B + c2.B}
}

// Multiply updates the color object the product of corresponding channels.
func (c Color) Multiply(c2 Color) Color {
	// c.R = c.R * c2.R
	// c.G = c.G * c2.G
	// c.B = c.B * c2.B
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
