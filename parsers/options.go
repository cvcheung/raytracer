package parsers

import (
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"

	"github.com/gonum/matrix/mat64"
)

// Options ...
type Options struct {
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

// WithOptions returns an options struct with the specified parameters
func WithOptions(parameters ...func(*Options)) *Options {
	origin := primitives.NewVec3(0, 0, 10)
	lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	distToFocus := 1.0
	aperature := 0.0
	camera := base.NewCameraFOV(origin, lookat, vertical, 20,
		float64(500)/float64(500), aperature, distToFocus, 0, 1)
	camera.ToggleBlur()
	opts := &Options{nx: 500, ny: 500, ns: 8, film: base.NewFilm(500, 500),
		camera: camera, ambientLight: materials.NewAmbientLight(textures.Black),
		lights: make([]materials.Light, 0, 10), world: objects.NewEmptyObjectList(10),
		mat: materials.NewBlinnphong(textures.White, textures.White, textures.White,
			textures.White, 1.0, materials.NewAmbientLight(textures.White)),
		transforms: make([]*mat64.Dense, 0, 3)}
	for _, f := range parameters {
		f(opts)
	}
	return opts
}

// SetDimensions ...
func (o *Options) SetDimensions(nx, ny int) {
	o.nx = nx
	o.ny = ny
	o.film = base.NewFilm(nx, ny)
}

// SetAntialiasing ...
func (o *Options) SetAntialiasing(ns int) {
	o.ns = ns
}

// SetCamera ...
func (o *Options) SetCamera(cam *base.Camera) {
	o.camera = cam
}

// SetAmbientLight ...
func (o *Options) SetAmbientLight(light materials.Light) {
	o.ambientLight = light
}

// SetMat ...
func (o *Options) SetMat(mat materials.Material) {
	o.mat = mat
}

// SetVFOV ...
func (o *Options) SetVFOV(vfov float64) {
	o.vfov = vfov
}

// SetAperture ...
func (o *Options) SetAperture(aperture float64) {
	o.aperture = aperture
}

// SetDistFocus ...
func (o *Options) SetDistFocus(distFocus float64) {
	o.distFocus = distFocus
}

// SetFOVCam ...
func (o *Options) SetFOVCam(fovcam bool) {
	o.fovcam = fovcam
}

// AddLights ...
func (o *Options) AddLights(lights ...materials.Light) {
	for _, l := range lights {
		o.lights = append(o.lights, l)
	}
}

// AddObjects ...
func (o *Options) AddObjects(objects ...objects.Object) {
	for _, obj := range objects {
		o.world.Add(obj)
	}
}

// GetCamera ...
func (o *Options) GetCamera() *base.Camera {
	return o.camera
}

// GetFilm ...
func (o *Options) GetFilm() *base.Film {
	return o.film
}

// GetWorld ...
func (o *Options) GetWorld() *objects.ObjectList {
	return o.world
}

// GetLights ...
func (o *Options) GetLights() []materials.Light {
	return o.lights
}

// GetAntialiasing ...
func (o *Options) GetAntialiasing() int {
	return o.ns
}
