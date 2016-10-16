package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
)

// Triangle ...
type Triangle struct {
	v1, v2, v3 primitives.Vec3
	t1, t2, t3 primitives.Vec3
	n1, n2, n3 primitives.Vec3
	mat        materials.Material
}

// NewTriangle ....
func NewTriangle(v1, v2, v3, t1, t2, t3 primitives.Vec3, mat materials.Material) *Triangle {
	triangle := Triangle{v1: v1, v2: v2, t1: t1, t2: t2, t3: t3, mat: mat}
	triangle.Normalize()
	return &triangle
}

// Hit returns true if a ray intersects with the triangle and stores the result
// in the passed record.
func (t *Triangle) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	return false
}

// Normalize sets the normals of the triangle given its vertices.
func (t *Triangle) Normalize() {
	e1 := t.v2.Subtract(t.v1)
	e2 := t.v3.Subtract(t.v1)
	normal := e1.Cross(e2).Normalize()
	t.n1 = normal
	t.n2 = normal
	t.n3 = normal
}

// BoundingBox returns the AABB for a sphere.
func (t *Triangle) BoundingBox(t0, t1 float64) (bool, *AABB) {
	minX := math.Min(math.Min(t.v1.X(), t.v2.X()), t.v3.X())
	minY := math.Min(math.Min(t.v1.Y(), t.v2.Y()), t.v3.Y())
	minZ := math.Min(math.Min(t.v1.Z(), t.v2.Z()), t.v3.Z())
	maxX := math.Max(math.Max(t.v1.X(), t.v2.X()), t.v3.X())
	maxY := math.Max(math.Max(t.v1.Y(), t.v2.Y()), t.v3.Y())
	maxZ := math.Max(math.Max(t.v1.Z(), t.v2.Z()), t.v3.Z())
	small := primitives.NewVec3(minX, minY, minZ)
	big := primitives.NewVec3(maxX, maxY, maxZ)
	return true, NewAABB(small, big)
}
