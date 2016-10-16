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
	phong        float64
	ambientLight Light
}

// NewBlinnphong returns a new material definition. Fuzz is bounded to less than or
// equal to 1.
func NewBlinnphong(ambient, diffuse, specular textures.Color, phong float64, ambientLight Light) Blinnphong {
	return Blinnphong{ambient, diffuse, specular, phong, ambientLight}
}

// Scatter calculates the incidental reflected ray if there is a reflection.
func (b Blinnphong) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, lights []Light) (bool, *primitives.Ray) {
	attenuation.Update(b.shade(rayIn, rec, depth, lights))
	reflected := rayIn.Direction().Normalize().Reflect(rec.Normal())
	scattered := primitives.NewRay(rec.Point(), reflected)
	rec.SetSpecular(b.specular)
	return scattered.Direction().Dot(rec.normal) > 0 && b.specular.NotBlack(), scattered
}

// Emitted is defined to implement the material interface.
func (b Blinnphong) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}

func (b Blinnphong) shade(rayIn *primitives.Ray, rec *HitRecord, depth int, lights []Light) textures.Color {
	color := textures.Black
	if depth == 0 {
		color = color.Add(b.ambient.Multiply(b.ambientLight.Intensity()))
	}
	for _, light := range lights {
		l := light.LVec(rec.Point())
		v := rayIn.Origin().Subtract(rec.Point()).Normalize()
		h := rec.normal.MultiplyScalar(2 * l.Dot(rec.normal)).Subtract(l)
		color = color.Add(b.diffuse.Multiply(light.Intensity()).
			MultiplyScalar(math.Max(0, rec.normal.Dot(l))))
		color = color.Add(b.specular.Multiply(light.Intensity()).
			MultiplyScalar(math.Max(0, math.Pow(v.Dot(h), b.phong))))
	}
	return color
}
