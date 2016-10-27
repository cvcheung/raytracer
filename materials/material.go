package materials

import (
	"raytracer/primitives"
	"raytracer/textures"
)

// Material is a container for how our polygons interact with light. All
// materials in the package must implement this interface.
type Material interface {
	Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, light Light, shadow bool) (bool, *primitives.Ray)
	Emitted(u, v float64, p primitives.Vec3) textures.Color
	GetAmbient() textures.Color
}
