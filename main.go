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

func shade(r *base.Ray, obj base.Object, depth int) base.Color {
	var rec base.HitRecord
	if obj.Hit(r, 0.001, MaxFloat64, &rec) {
		if depth < 50 {
			m := rec.Material()
			if bounce, scattered := m.Scatter(r, &rec); bounce {
				return m.Color().Multiply(shade(scattered, obj, depth+1))
			}
		}
		return base.Black
	}
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return base.White.MultiplyScalar(1.0 - t).Add(base.Blue.MultiplyScalar(t))
}

func main() {
	fp, err := os.Create("./output/part9-default.ppm")
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
	// lowerLeftCorner := base.NewVec3(-2.0, -1.0, -1.0)
	// horizontal := base.NewVec3(4.0, 0.0, 0.0)
	// vertical := base.NewVec3(0.0, 2.0, 0.0)
	// origin := base.NewVec3(0.0, 0.0, 0.0)
	// camera := base.NewCamera(lowerLeftCorner, horizontal, vertical, origin)
	origin := base.NewVec3(3, 3, 2)
	lookat := base.NewVec3(0.0, 0.0, -1.0)
	vertical := base.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 1.0
	aperature := 2.0
	camera := base.NewCameraFOV(origin, lookat, vertical, 20, float64(nx)/float64(ny), aperature, distToFocus)

	// Objects
	s1 := base.NewSphere(base.NewVec3(0, 0, -1), 0.5, base.NewLambertian(base.NewColor(0.1, 0.2, 0.5)))
	s2 := base.NewSphere(base.NewVec3(0, -100.5, -1), 100, base.NewLambertian(base.NewColor(0.8, 0.8, 0.0)))
	s3 := base.NewSphere(base.NewVec3(1, 0, -1), 0.5, base.NewMetal(base.NewColor(0.8, 0.6, 0.2), 1.0))
	s4 := base.NewSphere(base.NewVec3(-1, 0, -1), 0.5, base.NewDielectric(1.5))
	s5 := base.NewSphere(base.NewVec3(-1, 0, -1), -0.45, base.NewDielectric(1.5))

	objects := base.NewObjectList(5, s1, s2, s3, s4, s5)
	// R := math.Cos(math.Pi / 4)
	// s1 := base.NewSphere(base.NewVec3(-R, 0, -1), R, base.NewLambertian(base.NewColor(0, 0, 1)))
	// s2 := base.NewSphere(base.NewVec3(R, 0, -1), R, base.NewLambertian(base.NewColor(1, 0, 0)))
	// objects := base.NewObjectList(2, s1, s2)

	fp.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := base.NewEmptyColor()
			for k := 0; k < ns; k++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := camera.GetRay(u, v)
				color = color.Add(shade(r, objects, 0))
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
