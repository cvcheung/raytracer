package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"raytracer/base"
	"raytracer/primitives"
)

func main() {
	fp, err := os.Create("./output/refactor-2.ppm")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fp.Close()

	// Pixel counts
	nx := 1000
	ny := 500
	ns := 10

	// World space
	origin := primitives.NewVec3(13, 2, 3)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 10.0
	aperature := 0.1
	camera := base.NewCameraFOV(origin, lookat, vertical, 60,
		float64(nx)/float64(ny), aperature, distToFocus)
	// camera.ToggleBlur()

	// Objects
	objects := randomScene()

	fp.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := primitives.NewEmptyColor()
			for k := 0; k < ns; k++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := camera.GetRay(u, v)
				color = color.Add(shade(r, objects, 0))
			}
			color = color.DivideScalar(float64(ns))
			// Gamma correction
			color = primitives.NewColor(math.Sqrt(color.R),
				math.Sqrt(color.G),
				math.Sqrt(color.B))
			ir := int(255 * color.R)
			ig := int(255 * color.G)
			ib := int(255 * color.B)
			fp.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
}
