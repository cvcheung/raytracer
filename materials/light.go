package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
)

// Light ...
type Light interface {
	LVec(point primitives.Vec3) primitives.Vec3
	Intensity() textures.Color
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

// Intensity ...
func (a *AmbientLight) Intensity() textures.Color {
	return a.color
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
	return d.location
}

// Intensity ...
func (d *DirectionalLight) Intensity() textures.Color {
	return d.color
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

// Intensity ...
func (p *PointLight) Intensity() textures.Color {
	return p.color
}
