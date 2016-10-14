package base

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Film is the screen space that our world is rendered on.
type Film struct {
	width, height int
	Rect          image.Rectangle
	im            [][]color.Color
}

// NewFilm returns a new scene to be raytraced.
func NewFilm(width, height int) *Film {
	im := make([][]color.Color, width)
	for i := range im {
		im[i] = make([]color.Color, height)
	}
	return &Film{width, height, image.Rect(0, 0, width, height), im}
}

// Width returns the number of horizontal pixels in the film.
func (s *Film) Width() int {
	return s.width
}

// Height returns the number of vertical pixels in the film.
func (s *Film) Height() int {
	return s.height
}

// ColorModel returns the Image's color model.
func (s *Film) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (s *Film) Bounds() image.Rectangle {
	return s.Rect
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (s *Film) At(x, y int) color.Color {
	return s.im[x][y]
}

// Set updates the scene with the corresponding RGB values at the given pixel.
func (s *Film) Set(x, y int, r, g, b byte) {
	s.im[x][s.height-1-y] = color.RGBA{r, g, b, 255}
}

// Save the file the scene to disk with the corresponding filename in png format.
func (s *Film) Save(filename string) {
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", os.ModePerm)
	}

	fp, err := os.Create("./output/" + filename + ".png")
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
