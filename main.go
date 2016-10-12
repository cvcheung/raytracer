package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"raytracer/base"
)

// MaxFloat64 ...
const MaxFloat64 = 1.797693134862315708145274237317043567981e+308

func shade(r *base.Ray, obj base.Object) base.Color {
	var rec base.HitRecord
	if obj.Hit(r, 0.001, MaxFloat64, &rec) {
		target := rec.Point().Add(rec.Normal()).Add(base.RandomInUnitSphere())
		return shade(
			base.NewRay(rec.Point(),
				target.Subtract(rec.Point())), obj).
			MultiplyScalar(0.5)
	}
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return base.White.MultiplyScalar(1.0 - t).Add(base.Blue.MultiplyScalar(t))
}

func main() {
	fp, err := os.Create("./part5-shading-improved")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fp.Close()

	// Pixel counts
	nx := 1000
	ny := 500
	ns := 500

	// World space
	lowerLeftCorner := base.NewVec3(-2.0, -1.0, -1.0)
	horizontal := base.NewVec3(4.0, 0.0, 0.0)
	vertical := base.NewVec3(0.0, 2.0, 0.0)
	origin := base.NewVec3(0.0, 0.0, 0.0)
	camera := base.NewCamera(lowerLeftCorner, horizontal, vertical, origin)

	// Objects
	s1 := base.NewSphere(base.NewVec3(0, 0, -1), 0.5)
	s2 := base.NewSphere(base.NewVec3(0, -100.5, -1), 100)
	objects := base.NewObjectList(2, s1, s2)

	// Random number for anti-aliasing. We the set the seed to 7 because that is
	// what any sane person would do.
	rgen := rand.New(rand.NewSource(7))

	fp.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := base.NewEmptyColor()
			for k := 0; k < ns; k++ {
				u := (float64(i) + rgen.Float64()) / float64(nx)
				v := (float64(j) + rgen.Float64()) / float64(ny)
				r := camera.GetRay(u, v)
				color = color.Add(shade(r, objects))
			}
			color = color.DivideScalar(float64(ns))
			// Gamma correction
			color = base.NewColor(math.Sqrt(color.R),
				math.Sqrt(color.G),
				math.Sqrt(color.B))
			ir := int(255 * color.R)
			ig := int(255 * color.G)
			ib := int(255 * color.B)
			fp.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
}
