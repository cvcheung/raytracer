package main

import (
	"fmt"
	"os"
	"raytracer/base"
)

// MaxFloat64 ...
const MaxFloat64 = 1.797693134862315708145274237317043567981e+308

func shade(r *base.Ray, obj base.Object) base.Color {
	var rec base.HitRecord
	if obj.Hit(r, 0, MaxFloat64, &rec) {
		return base.NewColor(rec.Normal().X()+1, rec.Normal().Y()+1, rec.Normal().Z()+1).MultiplyScalar(0.5)
	}
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return base.White.MultiplyScalar(1.0 - t).Add(base.Blue.MultiplyScalar(t))
}

func main() {
	fp, err := os.Create("./part3")
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
	s1 := base.NewSphere(base.NewVec3(0, 0, -1), 0.5)
	s2 := base.NewSphere(base.NewVec3(0, -100.5, -1), 100)
	objects := base.NewObjectList(2, s1, s2)

	fp.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := base.NewRay(origin, lowerLeftCorner.Add(horizontal.MultiplyScalar(u)).Add(vertical.MultiplyScalar(v)))

			color := shade(r, objects)
			ir := int(255 * color.R)
			ig := int(255 * color.G)
			ib := int(255 * color.B)
			fp.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
}
