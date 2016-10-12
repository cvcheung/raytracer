package base

import "math"

// Sphere ...
type Sphere struct {
	center Vec3
	radius float64
}

// NewSphere constructs a new sphere object with the specified parameters.
func NewSphere(center Vec3, radius float64) Sphere {
	return Sphere{center, radius}
}

// Hit returns the value t the ray intersects with a point on the sphere,
// otherwise returns -1.0.
func (s Sphere) Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin().Subtract(s.center)
	a := r.Direction().Dot(r.Direction())
	b := 2 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		if t := (-b - math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			rec.t = t
			rec.p = r.PointAt(t)
			rec.normal = rec.p.Subtract(s.center).DivideScalar(s.radius)
			return true
		}
		if t := (-b + math.Sqrt(discriminant)) / (2 * a); t > tMin && t < tMax {
			rec.t = t
			rec.p = r.PointAt(t)
			rec.normal = rec.p.Subtract(s.center).DivideScalar(s.radius)
			return true
		}
	}
	return false
}

// Shade linearly blends white and blue.
func (s Sphere) Shade(r Ray) Color {
	// if t := s.Hit(r); t > 0 {
	// 	n := r.PointAt(t).Subtract(NewVec3(0, 0, -1)).Normalize()
	// 	return Color{n.x + 1, n.y + 1, n.z + 1}.MultiplyScalar(0.5)
	// }
	unitDirection := r.Direction().Normalize()
	t := 0.5 * (unitDirection.y + 1.0)
	return White.MultiplyScalar(1.0 - t).Add(Blue.MultiplyScalar(t))
}
