package main

import (
	"fmt"
	"os"
	"raytracer/base"
)

func main() {
	fp, err := os.Create("./part2-sphere")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fp.Close()

	// Pixel counts
	nx := 1000
	ny := 500

	// World space
	lowerLeftCorner := base.NewVec3(-2.0, -1.0, -1.0)
	horizontal := base.NewVec3(4.0, 0.0, 0.0)
	vertical := base.NewVec3(0.0, 2.0, 0.0)
	origin := base.NewVec3(0.0, 0.0, 0.0)

	// Objects
	sphere := base.NewSphere(base.NewVec3(0, 0, -1), 0.5)

	fp.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := base.NewRay(origin, lowerLeftCorner.Add(horizontal.MultiplyScalar(u)).Add(vertical.MultiplyScalar(v)))

			color := sphere.Shade(r)
			ir := int(255 * color.R)
			ig := int(255 * color.G)
			ib := int(255 * color.B)
			fp.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
}
