package objects

import (
	"raytracer/materials"
	"raytracer/primitives"
)

// Rectangle ...
type Rectangle struct {
	x0, x1, y0, y1, o float64
	mat               materials.Material
}

// NewRectangle returns a new rectangle with formed by the following four
// coordinates. o specifices the position on the z-axis. Why o - uvo.
func NewRectangle(x0, x1, y0, y1, o float64, mat materials.Material) Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle{x0, x1, y0, y1, o, mat}
}

// Hit ...
func (rect Rectangle) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	t := (rect.o - r.Origin().Z()) / r.Direction().Z()
	if t < tMin || t > tMax {
		return false
	}
	x := r.Origin().X() + t*r.Direction().X()
	y := r.Origin().Y() + t*r.Direction().Y()
	if x < rect.x0 || x > rect.x1 || y < rect.y0 || y > rect.y1 {
		return false
	}
	u := (x - rect.x0) / (rect.x1 - rect.x0)
	v := (y - rect.y0) / (rect.y1 - rect.y0)
	p := r.PointAt(t)
	rec.UpdateRecord(t, u, v, p, primitives.UnitZ, rect.mat)
	return true
}
