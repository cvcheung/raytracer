package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
)

// Light ...
type Light interface {
	LVec(point primitives.Vec3) primitives.Vec3
	Direction(point primitives.Vec3) primitives.Vec3
	Intensity() textures.Color
	Falloff() int
}

// AmbientLight ...
type AmbientLight struct {
	color textures.Color
}

// NewAmbientLight ...
func NewAmbientLight(color textures.Color) *AmbientLight {
	return &AmbientLight{color}
}

// LVec ...
func (a *AmbientLight) LVec(point primitives.Vec3) primitives.Vec3 {
	return primitives.Vec3{}
}

// Direction ...
func (a *AmbientLight) Direction(point primitives.Vec3) primitives.Vec3 {
	return primitives.Vec3{}
}

// Intensity ...
func (a *AmbientLight) Intensity() textures.Color {
	return a.color
}

// Falloff ...
func (a *AmbientLight) Falloff() int {
	return 0
}

// DirectionalLight is a our container for directional lighting.
type DirectionalLight struct {
	location primitives.Vec3
	color    textures.Color
}

// NewDirectionalLight ...
func NewDirectionalLight(location primitives.Vec3, color textures.Color) *DirectionalLight {
	return &DirectionalLight{location.Normalize(), color}
}

// LVec ...
func (d *DirectionalLight) LVec(point primitives.Vec3) primitives.Vec3 {
	return d.location.MultiplyScalar(-1)
}

// Direction ...
func (d *DirectionalLight) Direction(point primitives.Vec3) primitives.Vec3 {
	return d.location.MultiplyScalar(-1)
}

// Intensity ...
func (d *DirectionalLight) Intensity() textures.Color {
	return d.color
}

// Falloff ...
func (d *DirectionalLight) Falloff() int {
	return 0
}

// PointLight is a our container for point lighting.
type PointLight struct {
	location primitives.Vec3
	color    textures.Color
	falloff  int
}

// NewPointLight ...
func NewPointLight(location primitives.Vec3, color textures.Color, falloff int) *PointLight {
	return &PointLight{location, color, falloff}
}

// LVec ...
func (p *PointLight) LVec(point primitives.Vec3) primitives.Vec3 {
	return p.location.Subtract(point).Normalize()
}

// Direction ...
func (p *PointLight) Direction(point primitives.Vec3) primitives.Vec3 {
	return p.location.Subtract(point)
}

// Intensity ...
func (p *PointLight) Intensity() textures.Color {
	return p.color
}

// Falloff ...
func (p *PointLight) Falloff() int {
	return p.falloff
}
