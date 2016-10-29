package parsers

// Taken from https://github.com/angus-g/go-obj/blob/master/obj/obj.go

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parse a vertex (normal) string, a list of whitespace-separated
// floating point numbers.
func parseVertex(t []string) []float64 {
	x, _ := strconv.ParseFloat(t[0], 32)
	y, _ := strconv.ParseFloat(t[1], 32)
	z, _ := strconv.ParseFloat(t[2], 32)

	return []float64{float64(x), float64(y), float64(z)}
}

// Parse an element string, a list of whitespace-separated elements.
// Elements are of the form "<vi>/<ti>/<ni>" where indices are the
// vertex, texture coordinate and normal, respectively.
func parseElement(t []string) [][3]int32 {
	e := make([][3]int32, len(t))
	for i := 0; i < len(t); i++ {
		f := strings.Split(t[i], "/")
		for j := 0; j < len(f); j++ {
			// for now, just grab the vertex index
			if x, err := strconv.ParseInt(f[j], 10, 32); err == nil {
				e[i][j] = int32(x) - 1 // convert to 0-indexing
			} else {
				e[i][j] = -1
			}
		}
	}

	// convert quads to triangles
	if len(t) > 3 {
		e = append(e, e[0], e[2])
	}

	return e
}

// ParseObj ...
func ParseObj(filename string) ([]float64, []float64) {
	fp, _ := os.Open(filename)
	scanner := bufio.NewScanner(fp)

	vertices := [][]float64{}
	normals := [][]float64{}
	elements := [][3]int32{}

	vertOut := []float64{}
	normOut := []float64{}

	for scanner.Scan() {
		toks := strings.Fields(strings.TrimSpace(scanner.Text()))
		switch toks[0] {
		case "v":
			vertices = append(vertices, parseVertex(toks[1:]))
		case "vn":
			normals = append(normals, parseVertex(toks[1:]))
		case "f":
			elements = append(elements, parseElement(toks[1:])...)
		}
	}

	for _, e := range elements {
		if e[0] >= 0 {
			vertOut = append(vertOut, vertices[e[0]]...)
		}
		if e[2] >= 0 {
			normOut = append(normOut, normals[e[2]]...)
		}
	}

	return vertOut, normOut
}
