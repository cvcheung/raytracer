package materials

import (
	"math"
	"math/rand"
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/utils"
)

// Dielectric defines the dielectric material specified by the reflection index.
type Dielectric struct {
	reflectIdx float64
}

// NewDielectric returns a new diaelectric definition. 1.5 is used for glass.
func NewDielectric(reflectIdx float64) Dielectric {
	return Dielectric{reflectIdx}
}

// Scatter calculates the incidental reflected/refracted ray if there is a
// reflection/refraction.
func (d Dielectric) Scatter(rayIn *primitives.Ray, attenuation *textures.Color, rec *HitRecord, depth int, lights []Light) (bool, *primitives.Ray) {
	var outwardNormal primitives.Vec3
	var niOverNt, cosine, refractProb float64
	attenuation.Update(textures.White)
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
	refracted, refVec := rayIn.Direction().Refract(outwardNormal, niOverNt)
	if refracted {
		refractProb = utils.Schlick(cosine, d.reflectIdx)
	} else {
		refractProb = 1.0
	}
	if rand.Float64() < refractProb {
		reflected := rayIn.Direction().Reflect(rec.Normal())
		return true, primitives.NewRay(rec.Point(), reflected)
	}
	return true, primitives.NewRay(rec.Point(), refVec)
}

// Emitted is defined to implement the material interface.
func (d Dielectric) Emitted(u, v float64, p primitives.Vec3) textures.Color {
	return textures.Black
}
