package primitives

import "fmt"

// Ray ...
type Ray struct {
	origin, direction Vec3
	time              float64
}

// NewRay ...
func NewRay(origin, direction Vec3, options ...func(*Ray)) *Ray {
	r := &Ray{origin: origin, direction: direction}
	for _, f := range options {
		f(r)
	}
	return r
}

// WithTime is an optional parameter when generating a new ray.
func WithTime(time float64) func(*Ray) {
	return func(r *Ray) {
		r.time = time
	}
}

// Update modifies the ray with new parameters.
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

// Time returns the time of the ray to be used for motion blur.
func (r *Ray) Time() float64 {
	return r.time
}

// PointAt is our method to represent a ray as a paramtric function.
// Returns a point as a Vec3.
func (r *Ray) PointAt(f float64) Vec3 {
	return r.origin.Add(r.direction.MultiplyScalar(f))
}

func (r *Ray) String() string {
	return fmt.Sprintf("o: %s, d: %s", r.origin.String(), r.direction.String())
}
