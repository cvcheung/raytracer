package main

import (
	"math"
	"math/rand"
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"
	"runtime"
	"sync"
)

// TODO add func options to have a variety of backgrounds.
func shade(r *primitives.Ray, obj objects.Object, depth int) textures.Color {
	var rec materials.HitRecord
	if obj.Hit(r, 0.001, math.MaxFloat64, &rec) {
		m := rec.Material()
		emit := m.Emitted(rec.U(), rec.V(), rec.Point())
		if depth < 50 {
			var attenuation textures.Color
			if bounce, scattered := m.Scatter(r, &attenuation, &rec); bounce {
				return emit.Add(attenuation.Multiply(shade(scattered, obj, depth+1)))
			}
		}
		return emit
	}

	// Background color gradient. We use this to simulate light coming from the sky.
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return textures.Gradient(t)
	// return textures.Black
}

// TODO move rendering from main to here.
func render(ns int, fileName string, world objects.Object, camera *base.Camera, film *base.Film) {
	// Parallelization
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()

	for cpu := 0; cpu < numCPU; cpu++ {
		wg.Add(1)
		go func(column int) {
			for j := column; j < film.Height(); j += numCPU {
				for i := 0; i < film.Width(); i++ {
					color := textures.NewEmptyColor()
					for k := 0; k < ns*ns; k++ {
						var u, v float64
						if ns == 1 {
							u = (float64(i) + 0.5) / float64(film.Width())
							v = (float64(j) + 0.5) / float64(film.Height())
						} else {
							u = (float64(i) + rand.Float64()) / float64(film.Width())
							v = (float64(j) + rand.Float64()) / float64(film.Height())
						}
						r := camera.GetRay(u, v)
						color = color.Add(shade(r, world, 0))
					}
					color = color.DivideScalar(float64(ns * ns))
					// Gamma correction
					color = textures.NewColor(math.Sqrt(color.R),
						math.Sqrt(color.G),
						math.Sqrt(color.B))
					ir := byte(255 * color.R)
					ig := byte(255 * color.G)
					ib := byte(255 * color.B)
					film.Set(i, j, ir, ig, ib)
				}
			}
			wg.Done()
		}(cpu)
	}

	wg.Wait()
	film.Save(fileName)
}

// Generates the random scene from `Ray Tracing in One Weekend`
func randomScene() objects.Object {
	objList := objects.NewEmptyObjectList(500)
	// checkered := textures.NewCheckered(textures.NewColor(0.2, 0.3, 0.1),
	// 	textures.NewColor(0.9, 0.9, 0.9))
	// objList.Add(objects.NewSphere(primitives.NewVec3(0, -1000, 0), 1000,
	// 	materials.NewLambertian(checkered)))
	objList.Add(objects.NewSphere(primitives.NewVec3(0, -1000, 0), 1000,
		materials.NewLambertian(textures.NewColor(0.5, 0.5, 0.5))))
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {

			rndMat := rand.Float64()
			center := primitives.NewVec3(float64(a)+.9*rand.Float64(), 0.2,
				float64(b)+.9*rand.Float64())

			if center.Subtract(primitives.NewVec3(4, 0.2, 0)).Magnitude() > 0.9 {
				if rndMat < 0.8 {
					if move := rand.Float64(); move < 0.5 {
						objList.Add(objects.NewMovingSphere(center,
							center.Add(primitives.NewVec3(0, 0.5*(1+rand.Float64()), 0)),
							0, 1, 0.2, materials.NewLambertian(textures.NewRandomColor())))
					} else {
						objList.Add(objects.NewSphere(center, 0.2,
							materials.NewLambertian(textures.NewRandomColor())))
					}
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
		materials.NewLambertian(textures.NewColor(0.4, 0.2, 0.1))))
	objList.Add(objects.NewSphere(primitives.NewVec3(4, 1, 0), 1,
		materials.NewMetal(textures.NewColor(0.7, 0.6, 0.5), 0)))
	// return objList
	list := objList.List()
	return objects.NewBVHNode(list, len(list), 0, 1)
}
