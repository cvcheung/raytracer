package main

import (
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Pixel counts
	nx := 1000
	ny := 500
	ns := 1 // 8x Antialiasing
	film := base.NewFilm(nx, ny)

	// World space
	// origin := primitives.NewVec3(0, 0, 10)
	// // origin := primitives.NewVec3(4, 4, 4)
	// lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	// vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	// distToFocus := 1.0
	// aperature := 0.0
	// camera := base.NewCameraFOV(origin, lookat, vertical, 20,
	// 	float64(nx)/float64(ny), aperature, distToFocus, 0, 1)
	// camera.ToggleBlur()
	//
	eye := primitives.NewVec3(0, 0, 1)
	LL := primitives.NewVec3(-2, -1, 0)
	LR := primitives.NewVec3(2, -1, 0)
	UL := primitives.NewVec3(-2, 1, 0)
	UR := primitives.NewVec3(2, 1, 0)
	camera := base.NewCameraFromCoordinates(LL, LR, UL, UR, eye, float64(nx), float64(ny))

	// Objects
	// world := randomScene()
	ambient := textures.NewColor(0, 0, 0)
	diffuse := textures.NewColor(0.5, 0.5, 0.5)
	specular := textures.NewColor(0.5, 0.5, 0.5)
	phong := 1.0
	ambientLight := materials.NewAmbientLight(textures.NewColor(0.5, 0.5, 0.5))
	world := objects.NewSphere(primitives.NewVec3(0, 0, -1), 1, materials.NewBlinnphong(ambient, diffuse, specular, phong, ambientLight))
	lights := []materials.Light{materials.NewDirectionalLight(primitives.NewVec3(-1, 1, 1), textures.NewColor(.35, .7, 1))}

	scene := base.NewScene(camera, film, world, lights, ns)
	scene.Render("test")

}
