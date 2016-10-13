package materials

import (
	"math"
	"math/rand"
	"raytracer/primitives"
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

// Color always returns white since glass is clear.
func (d Dielectric) Color() primitives.Color {
	return primitives.White
}

// Scatter calculates the incidental reflected/refracted ray if there is a
// reflection/refraction.
func (d Dielectric) Scatter(rayIn *primitives.Ray, rec *HitRecord) (bool, *primitives.Ray) {
	var outwardNormal primitives.Vec3
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
	// refracted, refVec := rayIn.Direction().Refract(outwardNormal, niOverNt)
	refracted, refVec := utils.Refract(rayIn.Direction(), outwardNormal, niOverNt)
	if refracted {
		refractProb = utils.Schlick(cosine, d.reflectIdx)
	} else {
		refractProb = 1.0
	}
	if rand.Float64() < refractProb {
		reflected := utils.Reflect(rayIn.Direction(), rec.Normal())
		// reflected := rayIn.Direction().Reflect(rec.Normal())
		return true, primitives.NewRay(rec.Point(), reflected)
	}
	return true, primitives.NewRay(rec.Point(), refVec)
}
