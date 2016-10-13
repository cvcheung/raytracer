package main

import (
	"math/rand"
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

// TODO move rendering from main to here.
func render() {

}

// Generates the random scene from `Ray Tracing in One Weekend`
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
				} else if rndMat < 0.95 {
					objList.Add(objects.NewSphere(center, 0.2,
						materials.NewRandomMetal()))
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
	objList.Add(objects.NewSphere(primitives.NewVec3(4, 1, 0), 1,
		materials.NewMetal(primitives.NewColor(0.7, 0.6, 0.5), 0)))
	return objList
}
