package base

// Ray ...
type Ray struct {
	origin, direction Vec3
}

// NewRay ...
func NewRay(origin, direction Vec3) *Ray {
	return &Ray{origin, direction}
}

// Origin returns the starting point of the ray represented as a Vec3.
func (r *Ray) Origin() Vec3 {
	return r.origin
}

// Direction returns the direction the ray is traveling in represented as a
// Vec3.
func (r *Ray) Direction() Vec3 {
	return r.direction
}

// PointAt is our method to represent a ray as a paramtric function.
// Returns a point as a Vec3.
func (r *Ray) PointAt(f float64) Vec3 {
	return r.origin.Add(r.direction.MultiplyScalar(f))
}

// Shade linearly blends white and blue.
func (r *Ray) Shade() Color {
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.y + 1.0)
	return White.MultiplyScalar(1.0 - t).Add(Blue.MultiplyScalar(t))
}
