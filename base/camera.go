package base

import (
	"math"
	"raytracer/primitives"
	"raytracer/utils"
)

// Camera is a the container for the information about the viewer.
type Camera struct {
	ll, horizontal, vertical, origin primitives.Vec3
	u, v, w                          primitives.Vec3
	lensRadius                       float64
	blur                             bool
}

// NewCamera returns a new camera object with the specified parameters.
func NewCamera(ll, horizontal, vertical, origin, u, v, w primitives.Vec3, lensRadius float64) *Camera {
	return &Camera{ll, horizontal, vertical, origin,
		u, v, w,
		lensRadius, false}
}

// NewCameraFOV returns a new camera object from a particular viewpoint with the
// specified FOV.
func NewCameraFOV(origin, lookat, vup primitives.Vec3, vfov, aspect, aperature, distToFocus float64) *Camera {
	lensRadius := aperature / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := origin.Subtract(lookat).Normalize()
	u := vup.Cross(w).Normalize()
	v := w.Cross(u)
	ll := origin.
		Subtract(u.MultiplyScalar(halfWidth * distToFocus)).
		Subtract(v.MultiplyScalar(halfHeight * distToFocus)).
		Subtract(w.MultiplyScalar(distToFocus))
	horizontal := u.MultiplyScalar(2 * halfWidth * distToFocus)
	vertical := v.MultiplyScalar(2 * halfHeight * distToFocus)
	return NewCamera(ll, horizontal, vertical, origin, u, v, w, lensRadius)
}

// ToggleBlur turns blur to on if off and vice versa.
func (c *Camera) ToggleBlur() bool {
	c.blur = !c.blur
	return c.blur
}

// GetRay returns a ray from the point of view of the camera.
func (c *Camera) GetRay(u, v float64) *primitives.Ray {
	if c.blur {
		rd := utils.RandomInUnitDisk().MultiplyScalar(c.lensRadius)
		offset := c.u.MultiplyScalar(rd.X()).Add(c.v.MultiplyScalar(rd.Y()))
		return primitives.NewRay(c.origin.Add(offset), c.ll.
			Add(c.horizontal.MultiplyScalar(u)).
			Add(c.vertical.MultiplyScalar(v)).
			Subtract(c.origin).Subtract(offset))
	}
	return primitives.NewRay(c.origin, c.ll.
		Add(c.horizontal.MultiplyScalar(u)).
		Add(c.vertical.MultiplyScalar(v)).
		Subtract(c.origin))
}
