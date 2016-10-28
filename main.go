package main

import (
	"flag"
	"raytracer/base"
	"raytracer/parsers"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	opts := parsers.WithOptions()

	x := flag.Uint("x", 500, "Specifies the width of the image.")
	y := flag.Uint("y", 500, "Specifies the height of the image.")
	aa := flag.Uint("aa", 8, "Sets the antialising amount.")
	input := flag.String("f", "", "File to load.")
	filename := flag.String("o", "output", "The filename.")
	random := flag.Bool("r", false, "Generate a random scene.")
	blur := flag.Bool("blur", false, "Turns on camera blur, effects change based on camera.")
	vfov := flag.Float64("vfov", 60, "Sets the camera fov, requires fovcam.")
	distFocus := flag.Float64("dist", 1, "Sets the distance to focus.")
	aperture := flag.Float64("apt", 0, "Sets the aperature of the camera, requires fovcam.")
	fovcam := flag.Bool("fovcam", false, "Use a camera with a specified field of view.")
	flag.Parse()

	if *random {

	}
	opts.SetVFOV(*vfov)
	opts.SetAperture(*aperture)
	opts.SetFOVCam(*fovcam)
	opts.SetDistFocus(*distFocus)
	opts.SetDimensions(int(*x), int(*y))
	opts.SetAntialiasing(int(*aa))
	parsers.ParseFile(*input, opts)

	// World space
	// origin := primitives.NewVec3(0, 0, 10)
	// lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	// vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	// distToFocus := 1.0
	// aperature := 0.0
	// camera := base.NewCameraFOV(origin, lookat, vertical, 20,
	// 	float64(nx)/float64(ny), aperature, distToFocus, 0, 1)
	// camera.ToggleBlur()

	if *blur {
		opts.GetCamera().ToggleBlur()
	}
	scene := base.NewScene(opts.GetCamera(), opts.GetFilm(), opts.GetWorld(),
		opts.GetLights(), opts.GetAntialiasing())
	scene.Render(*filename)
}
