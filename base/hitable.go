package base

// HitRecord is a simple record to that records the information regarding where
// the ray hit.
type HitRecord struct {
	t         float64
	p, normal Vec3
}

// Object defines base class for 3D objects.
type Object interface {
	Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool
}
