package main

import (
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"

	"github.com/gonum/matrix/mat64"
)

type options struct {
	nx, ny, ns                int
	vfov, aperture, distFocus float64
	film                      *base.Film
	camera                    *base.Camera
	ambientLight              materials.Light
	lights                    []materials.Light
	world                     *objects.ObjectList
	mat                       materials.Material
	transforms                []*mat64.Dense
	fovcam                    bool
}

// withOptions returns an options struct with the specified parameters
func withOptions(parameters ...func(*options)) *options {
	origin := primitives.NewVec3(0, 0, 10)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 1.0
	aperature := 0.0
	camera := base.NewCameraFOV(origin, lookat, vertical, 20,
		float64(500)/float64(500), aperature, distToFocus, 0, 1)
	camera.ToggleBlur()
	opts := &options{nx: 500, ny: 500, ns: 8, film: base.NewFilm(500, 500),
		camera: camera, ambientLight: materials.NewAmbientLight(textures.White),
		lights: make([]materials.Light, 0, 10), world: objects.NewEmptyObjectList(10),
		mat: materials.NewBlinnphong(textures.White, textures.White, textures.White,
			textures.White, 1.0, materials.NewAmbientLight(textures.White)),
		transforms: make([]*mat64.Dense, 0, 3)}
	for _, f := range parameters {
		f(opts)
	}
	return opts
}

func (o *options) setDimensions(nx, ny int) {
	o.nx = nx
	o.ny = ny
	o.film = base.NewFilm(nx, ny)
}

func (o *options) setAntialiasing(ns int) {
	o.ns = ns
}

func (o *options) setCamera(cam *base.Camera) {
	o.camera = cam
}

func (o *options) setAmbientLight(light materials.Light) {
	o.ambientLight = light
}

func (o *options) setMat(mat materials.Material) {
	o.mat = mat
}
func (o *options) addLights(lights ...materials.Light) {
	for _, l := range lights {
		o.lights = append(o.lights, l)
	}
}

func (o *options) addObjects(objects ...objects.Object) {
	for _, obj := range objects {
		o.world.Add(obj)
	}
}
