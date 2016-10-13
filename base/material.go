package base

import (
	"math"
	"math/rand"
)

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

// Refract returns the vector from refraction according to Snell's Law.
func Refract(v, n Vec3, niOverNt float64) (bool, Vec3) {
	uv := v.Normalize()
	dt := uv.Dot(n)
	discriminant := 1.0 - (niOverNt * niOverNt * (1 - dt*dt))
	if discriminant > 0 {
		refracted := uv.Subtract(n.MultiplyScalar(dt)).MultiplyScalar(niOverNt).
			Subtract(n.MultiplyScalar(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, Vec3{}
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
	fuzz   float64
}

// NewMetal ...
func NewMetal(color Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{color, fuzz}
}

// NewRandomMetal ...
func NewRandomMetal() Metal {
	color := NewColor(0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()))
	fuzz := 0.5 * rand.Float64()
	return Metal{color, fuzz}
}

// Color ...
func (m Metal) Color() Color {
	return m.albedo
}

// Scatter ...
func (m Metal) Scatter(rayIn *Ray, rec *HitRecord) (bool, *Ray) {
	reflected := Reflect(rayIn.Direction().Normalize(), rec.Normal())
	scattered := NewRay(rec.Point(),
		reflected.Add(RandomInUnitSphere().MultiplyScalar(m.fuzz)))
	return scattered.Direction().Dot(rec.normal) > 0, scattered
}

// Dielectric ...
type Dielectric struct {
	reflectIdx float64
}

// NewDielectric ...
func NewDielectric(reflectIdx float64) Dielectric {
	return Dielectric{reflectIdx}
}

// Color ...
func (d Dielectric) Color() Color {
	return White
}

// Scatter ...
func (d Dielectric) Scatter(rayIn *Ray, rec *HitRecord) (bool, *Ray) {
	var outwardNormal Vec3
	var niOverNt, cosine, refractProb float64
	if rayIn.Direction().Dot(rec.Normal()) > 0 {
		outwardNormal = rec.Normal().MultiplyScalar(-1)
		niOverNt = d.reflectIdx
		cosine = rayIn.Direction().Dot(rec.normal) / rayIn.Direction().Magnitude()
		cosine = math.Sqrt(1 - d.reflectIdx*d.reflectIdx*(1-cosine*cosine))
	} else {
		outwardNormal = rec.Normal()
		niOverNt = 1.0 / d.reflectIdx
		cosine = -rayIn.Direction().Dot(rec.normal) / rayIn.Direction().Magnitude()
	}
	refracted, refVec := Refract(rayIn.Direction(), outwardNormal, niOverNt)
	if refracted {
		refractProb = Schlick(cosine, d.reflectIdx)
	} else {
		refractProb = 1.0
	}
	if rand.Float64() < refractProb {
		reflected := Reflect(rayIn.Direction(), rec.Normal())
		return true, NewRay(rec.Point(), reflected)
	}
	return true, NewRay(rec.Point(), refVec)
}

// Schlick angle reflection approximation.
func Schlick(cosine, reflectIdx float64) float64 {
	r0 := (1 - reflectIdx) / (1 + reflectIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
