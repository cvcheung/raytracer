package base

// Material is a container for how are polygons interact with light.
type Material interface {
	Color() Color
	Scatter(rayIn *Ray, rec *HitRecord) (bool, *Ray)
}

// Reflect returns the vector in which the incidental vector reflects relative
// to the normal.
func Reflect(v, n Vec3) Vec3 {
	return v.Subtract(n.MultiplyScalar(2 * v.Dot(n)))
}

// Lambertian ...
type Lambertian struct {
	albedo Color
}

// NewLambertian ...
func NewLambertian(color Color) Lambertian {
	return Lambertian{color}
}

// Color ...
func (l Lambertian) Color() Color {
	return l.albedo
}

// Scatter ...
func (l Lambertian) Scatter(rayIn *Ray, rec *HitRecord) (bool, *Ray) {
	target := rec.Point().Add(rec.Normal()).Add(RandomInUnitSphere())
	return true, NewRay(rec.Point(), target.Subtract(rec.Point()))
}

// Metal ...
type Metal struct {
	albedo Color
}

// NewMetal ...
func NewMetal(color Color) Metal {
	return Metal{color}
}

// Color ...
func (m Metal) Color() Color {
	return m.albedo
}

// Scatter ...
func (m Metal) Scatter(rayIn *Ray, rec *HitRecord) (bool, *Ray) {
	reflected := Reflect(rayIn.Direction().Normalize(), rec.Normal())
	scattered := NewRay(rec.Point(), reflected)
	return scattered.Direction().Dot(rec.normal) > 0, scattered
}
