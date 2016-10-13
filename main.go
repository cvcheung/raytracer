package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
)

// MaxFloat64 ...
const MaxFloat64 = 1.797693134862315708145274237317043567981e+308

func shade(r *primitives.Ray, obj objects.Object, depth int) primitives.Color {
	var rec materials.HitRecord
	if obj.Hit(r, 0.001, MaxFloat64, &rec) {
		if depth < 50 {
			m := rec.Material()
			if bounce, scattered := m.Scatter(r, &rec); bounce {
				return m.Color().Multiply(shade(scattered, obj, depth+1))
			}
		}
		return primitives.Black
	}

	// Background color gradient
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return primitives.Gradient(t)
}

func randomScene() *objects.ObjectList {
	objList := objects.NewEmptyObjectList(500)
	objList.Add(objects.NewSphere(primitives.NewVec3(0, -1000, 0), 1000,
		materials.NewLambertian(primitives.NewColor(0.5, 0.5, 0.5))))
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {

			rndMat := rand.Float64()
			center := primitives.NewVec3(float64(a)+.9*rand.Float64(), 0.2,
				float64(b)+.9*rand.Float64())

			if center.Subtract(primitives.NewVec3(4, 0.2, 0)).Magnitude() > 0.9 {
				if rndMat < 0.8 {
					objList.Add(objects.NewSphere(center, 0.2,
						materials.NewLambertian(primitives.NewRandomColor())))
					// } else if rndMat < 0.95 {
					// 	objList.Add(objects.NewSphere(center, 0.2,
					// 		materials.NewRandomMetal()))
				} else {
					objList.Add(objects.NewSphere(center, 0.2,
						materials.NewDielectric(1.5)))
				}
			}
		}
	}
	objList.Add(objects.NewSphere(primitives.NewVec3(0, 1, 0), 1,
		materials.NewDielectric(1.5)))
	objList.Add(objects.NewSphere(primitives.NewVec3(-4, 1, 0), 1,
		materials.NewLambertian(primitives.NewColor(0.4, 0.2, 0.1))))
	// objList.Add(objects.NewSphere(primitives.NewVec3(4, 1, 0), 1,
	// 	materials.NewMetal(primitives.NewColor(0.7, 0.6, 0.5), 0)))
	return objList
}

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
	ns := 2

	// World space
	origin := primitives.NewVec3(13, 2, 3)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 10.0
	aperature := 0.1
	camera := base.NewCameraFOV(origin, lookat, vertical, 20,
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
