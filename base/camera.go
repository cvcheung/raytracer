package base

// Camera is a the container for the information about the viewer.
type Camera struct {
	ll, horizontal, vertical, origin Vec3
}

// NewCamera returns a new Camera object with the specified parameters.
func NewCamera(ll, horizontal, vertical, origin Vec3) *Camera {
	return &Camera{ll, horizontal, vertical, origin}
}

// GetRay returns a ray from the point of view of the camera.
func (c *Camera) GetRay(u, v float64) *Ray {
	return NewRay(c.origin, c.ll.
		Add(c.horizontal.MultiplyScalar(u)).
		Add(c.vertical.MultiplyScalar(v)).
		Subtract(c.origin))
}
