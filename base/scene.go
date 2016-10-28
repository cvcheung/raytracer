package base

import (
	"math"
	"math/rand"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"
	"runtime"
	"sync"
)

// Scene ...
type Scene struct {
	camera *Camera
	film   *Film
	world  objects.Object
	lights []materials.Light
	ns     int
}

// NewScene ...
func NewScene(camera *Camera, film *Film, world objects.Object, lights []materials.Light, ns int) *Scene {
	return &Scene{camera, film, world, lights, ns}
}

// TODO add func options to have a variety of backgrounds.
func (s *Scene) shade(r *primitives.Ray, obj objects.Object, depth int) textures.Color {
	var rec materials.HitRecord
	if obj.Hit(r, 0.001, math.MaxFloat64, &rec) {
		m := rec.Material()
		emit := m.Emitted(rec.U(), rec.V(), rec.Point())
		if depth < 50 {
			var attenuation textures.Color
			finalColor := textures.Black
			if depth == 0 {
				finalColor = finalColor.Add(m.GetAmbient())
			}
			for _, light := range s.lights {
				var shadowRec materials.HitRecord
				direction := light.Direction(rec.Point())
				shadowRay := primitives.NewRay(rec.Point(), direction)
				shadow := obj.Hit(shadowRay, 0.001, math.MaxFloat64, &shadowRec)
				if shadow {
					continue
				}
				m.Scatter(r, &attenuation, &rec, depth, light, shadow)
				finalColor = finalColor.Add(attenuation)
			}
			if bounce, scattered := m.Scatter(r, &attenuation, &rec, depth, nil, false); bounce {
				if rec.Reflective().NotBlack() {
					return emit.Add(finalColor).Add(rec.Reflective().Multiply(s.shade(scattered, obj, depth+1)))
				}
			}
			return emit.Add(finalColor)
		}
		return emit
	}

	return textures.Black
}

// Render ...
func (s *Scene) Render(fileName string, random bool) {
	// Parallelization
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()

	for cpu := 0; cpu < numCPU; cpu++ {
		wg.Add(1)
		go func(column int) {
			for j := column; j < s.film.Height(); j += numCPU {
				for i := 0; i < s.film.Width(); i++ {
					color := textures.NewEmptyColor()
					for k := 0; k < s.ns*s.ns; k++ {
						var u, v float64
						if s.ns == 1 {
							u = (float64(i) + 0.5) / float64(s.film.Width())
							v = (float64(j) + 0.5) / float64(s.film.Height())
						} else {
							u = (float64(i) + rand.Float64()) / float64(s.film.Width())
							v = (float64(j) + rand.Float64()) / float64(s.film.Height())
						}
						r := s.camera.GetRay(u, v)
						if random {
							color = color.Add(s.shadeRandom(r, s.world, 0))
						} else {
							color = color.Add(s.shade(r, s.world, 0))
						}
					}
					color = color.DivideScalar(float64(s.ns * s.ns))
					color = color.Clip()
					// Gamma correction
					color = textures.NewColor(math.Sqrt(color.R),
						math.Sqrt(color.G),
						math.Sqrt(color.B))
					ir := byte(255 * color.R)
					ig := byte(255 * color.G)
					ib := byte(255 * color.B)
					s.film.Set(i, j, ir, ig, ib)
				}
			}
			wg.Done()
		}(cpu)
	}

	wg.Wait()
	s.film.Save(fileName)
}

// Backup
// // TODO add func options to have a variety of backgrounds.
func (s *Scene) shadeRandom(r *primitives.Ray, obj objects.Object, depth int) textures.Color {
	var rec materials.HitRecord
	if obj.Hit(r, 0.001, math.MaxFloat64, &rec) {
		m := rec.Material()
		emit := m.Emitted(rec.U(), rec.V(), rec.Point())
		if depth < 50 {
			var attenuation textures.Color
			if bounce, scattered := m.Scatter(r, &attenuation, &rec, depth, nil, false); bounce {
				return emit.Add(attenuation.Multiply(s.shadeRandom(scattered, obj, depth+1)))
			}
			return emit.Add(attenuation)
		}
		return emit
	}

	// Background color gradient. We use this to simulate light coming from the sky.
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return textures.Gradient(t)
}
