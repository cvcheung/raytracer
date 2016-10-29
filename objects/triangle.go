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
func NewTriangle(v1, v2, v3 primitives.Vec3, mat materials.Material) *Triangle {
	triangle := Triangle{v1: v1, v2: v2, v3: v3, mat: mat}
	triangle.Normalize()
	return &triangle
}

func NewTriangleNormals(v1, v2, v3, n1, n2, n3 primitives.Vec3, mat materials.Material) *Triangle {
	triangle := Triangle{v1: v1, v2:v2, v3:v3, n1:n1, n2:n2, n3:n3, mat:mat}
	return &triangle
}

// Hit returns true if a ray intersects with the triangle and stores the result
// in the passed record.
func (t *Triangle) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	e1 := t.v2.Subtract(t.v1)
	e2 := t.v3.Subtract(t.v1)
	pv := r.Direction().Cross(e2)
	det := pv.Dot(e1)

	if det > -0.00001 && det < 0.00001 {
		return false
	}

	divDet := 1 / det
	eyeVec := r.Origin().Subtract(t.v1)
	u := eyeVec.Dot(pv) * divDet
	if u < 0 || u > 1 {
		return false
	}
	e2Vec := eyeVec.Cross(e1)
	v := r.Direction().Dot(e2Vec) * divDet
	if v < 0 || u+v > 1 {
		return false
	}
	// Record t
	inter := e2.Dot(e2Vec) * divDet
	if inter < tMin || inter > tMax {
		return false
	}
	p := r.PointAt(inter)
	gamma := (t.v1.X() - t.v3.X()) + (t.v3.Y()-t.v1.Y())*(t.v2.X()-t.v1.X())
	if gamma != 0 {
		gamma = ((p.Y()-t.v1.Y())*(t.v2.X()-t.v1.X()) + (p.X() - t.v1.X())) / gamma
	}
	beta := (p.X() - t.v1.X()) + gamma*(t.v1.X()-t.v3.X())
	if t.v2.X()-t.v1.X() != 0 {
		beta = beta / (t.v2.X() - t.v1.X())
	}
	U := t.n1.Subtract(t.n2)
	V := t.n1.Subtract(t.n3)
	normal := t.n1.Add(U.MultiplyScalar(beta)).Add(V.MultiplyScalar(gamma)).Normalize()
	rec.UpdateRecord(inter, u, v, p, normal, t.mat)
	return true
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
