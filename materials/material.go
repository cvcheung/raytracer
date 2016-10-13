package materials

import "raytracer/primitives"

// Material is a container for how our polygons interact with light. All
// materials in the package must implement this interface.
type Material interface {
	Color() primitives.Color
	Scatter(rayIn *primitives.Ray, rec *HitRecord) (bool, *primitives.Ray)
}
