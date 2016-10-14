package main

import (
	"raytracer/base"
	"raytracer/primitives"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Pixel counts
	nx := 1000
	ny := 500
	ns := 8 // 8x Antialiasing
	film := base.NewFilm(nx, ny)

	// World space
	origin := primitives.NewVec3(13, 2, 3)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 10.0
	aperature := 0.1
	camera := base.NewCameraFOV(origin, lookat, vertical, 20,
		float64(nx)/float64(ny), aperature, distToFocus)
	camera.ToggleBlur()

	// Objects
	world := randomScene()
	render(ns, "part13", world, camera, film)

}
