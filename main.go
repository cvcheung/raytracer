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
	// origin := primitives.NewVec3(4, 4, 4)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 10.0
	aperature := 0.0
	camera := base.NewCameraFOV(origin, lookat, vertical, 20,
		float64(nx)/float64(ny), aperature, distToFocus, 0, 1)
	camera.ToggleBlur()

	// Objects
	world := randomScene()
	// s1 := objects.NewSphere(primitives.NewVec3(0, -1000, 0), 1000, materials.NewLambertian(textures.NewColor(0.8, 0.8, 0)))
	// s2 := objects.NewSphere(primitives.NewVec3(0, 2, 0), 2, materials.NewLambertian(textures.NewColor(0.8, 0.8, 0)))
	// s3 := objects.NewSphere(primitives.NewVec3(0, 7, 0), 2, materials.NewDiffuseLight(textures.NewColor(4, 4, 4)))
	// s4 := objects.NewRectangle(3, 5, 1, 3, -2, materials.NewDiffuseLight(textures.NewColor(4, 4, 4)))
	// s5 := objects.NewSphere(primitives.NewVec3(-1, 0, -1), -0.45, materials.NewDielectric(1.5))

	// world := objects.NewObjectList(4, s1, s2, s3, s4)
	render(ns, "part15-5050", world, camera, film)

}
