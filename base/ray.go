package base

// Ray ...
type Ray struct {
	origin, direction Vec3
}

// NewRay ...
func NewRay(origin, direction Vec3) *Ray {
	return &Ray{origin, direction}
}

// Update modifies the ray with new parameters
func (r *Ray) Update(origin, direction Vec3) {
	r.origin = origin
	r.direction = direction
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
