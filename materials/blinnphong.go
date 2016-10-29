package materials

import (
	"math"
	"raytracer/primitives"
	"raytracer/textures"
)

// Blinnphong defines a generic material that does the Blinn-Phong Shading.
type Blinnphong struct {
	ambient      textures.Color
	diffuse      textures.Color
	specular     textures.Color
	reflective   textures.Color
	phong        float64
	ambientLight Light
}

// NewBlinnphong returns a new material definition. Fuzz is bounded to less than or
// equal to 1.
func NewBlinnphong(ambient, diffuse, specular, reflective textures.Color, phong float64, ambientLight Light) Blinnphong {
	return Blinnphong{ambient, diffuse, specular, reflective, phong, ambientLight}
}

// Scatter calculates the incidental reflected ray if there is a reflection.
func (b Blinnphong) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, light Light, shadow bool) (bool, *primitives.Ray) {
	if light != nil {
		attenuation.Update(b.shade(rayIn, rec, depth, light, shadow))
	}
	reflected := rayIn.Direction().Normalize().Reflect(rec.Normal())
	scattered := primitives.NewRay(rec.Point(), reflected)
	rec.SetReflective(b.reflective)
	return scattered.Direction().Dot(rec.normal) > 0 && b.reflective.NotBlack(), scattered
}

// Emitted is defined to implement the material interface.
func (b Blinnphong) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}

func (b Blinnphong) shade(rayIn *primitives.Ray, rec *HitRecord, depth int, light Light, shadow bool) textures.Color {
	color := textures.Black
	if !shadow {
		n := rec.normal
		l := light.LVec(rec.Point())
		intensity := light.Intensity()
		distance := light.Direction(rec.Point()).Magnitude()
		if light.Falloff() == 1 {
			intensity = intensity.DivideScalar(distance)
		} else if light.Falloff() == 2 {
			intensity = intensity.DivideScalar(distance * distance)
		}
		color = color.Add(b.diffuse.Multiply(intensity).
			MultiplyScalar(math.Max(0, n.Dot(l))))

		v := rayIn.Origin().Subtract(rec.Point()).Normalize()
		lh := l.MultiplyScalar(-1)
		rh := n.MultiplyScalar(2 * l.Dot(n))
		r := rh.Add(lh)

		if r.Dot(v)/(r.Magnitude()*v.Magnitude()) < 0 {
			return color
		}

		rbv := r.Dot(v)
		rbv = math.Pow(rbv, b.phong)

		color = color.Add(b.specular.Multiply(intensity).
			MultiplyScalar(rbv))

	}
	return color
}

// GetAmbient ...
func (b Blinnphong) GetAmbient() textures.Color {
	return b.ambient.Multiply(b.ambientLight.Intensity())
}
