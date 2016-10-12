package base

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Scene ...
type Scene struct {
	eye, UL, UR, LR, LL Vec3
	Rect                image.Rectangle
	image               [][]color.Color
}

// NewScene returns a new scene to be raytraced.
func NewScene(eye, UL, UR, LR, LL Vec3) Scene {
	return Scene{}
}

// ColorModel returns the Image's color model.
func (s *Scene) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (s *Scene) Bounds() image.Rectangle {
	return s.Rect
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (s *Scene) At(x, y int) color.Color {
	return s.image[x][y]
}

// Set updates the scene with the corresponding RGB values at the given pixel.
func (s *Scene) Set(x, y int, r, g, b byte) {
	s.image[x][y] = color.RGBA{r, g, b, 255}
}

// Save the file the scene to disk with the corresponding filename in png format.
func (s *Scene) Save(filename string) {
	fp, err := os.Create("./" + filename + ".png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fp.Close()

	err = png.Encode(fp, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
