package textures

import "raytracer/primitives"

// Texture defines base class for coloring of materials.
type Texture interface {
	GetColor(u, v float64, p primitives.Vec3) Color
}
