package transformations

import (
	"math"
	"raytracer/primitives"

	"github.com/gonum/matrix/mat64"
)

// NewTranslationMatrix ....
func NewTranslationMatrix(x, y, z float64) *mat64.Dense {
	matrix := make([]float64, 16)
	matrix[0] = 1
	matrix[5] = 1
	matrix[10] = 1
	matrix[3] = x
	matrix[7] = y
	matrix[11] = z
	matrix[15] = 1
	return mat64.NewDense(4, 4, matrix)
}

// NewRotationMatrix ...
func NewRotationMatrix(x, y, z float64) *mat64.Dense {
	matrix := make([]float64, 16)
	theta := primitives.NewVec3(x, y, z).Magnitude() * math.Pi / 180
	direction := primitives.NewVec3(x, y, z).Normalize()

	matrix[1] = -direction.Z()
	matrix[2] = direction.Y()
	matrix[4] = direction.Z()
	matrix[6] = -direction.X()
	matrix[8] = -direction.Y()
	matrix[9] = direction.X()
	matrix[15] = 1

	crossMatrix := mat64.NewDense(4, 4, matrix)
	crossSquared := mat64.NewDense(4, 4, nil)
	crossSquared.Mul(crossMatrix, crossMatrix)

	sine := mat64.NewDense(4, 4, nil)
	cosine := mat64.NewDense(4, 4, nil)

	sine.Scale(math.Sin(theta), crossMatrix)
	cosine.Scale(1-math.Cos(theta), crossSquared)

	iMatrix := make([]float64, 16)
	iMatrix[0] = 1
	iMatrix[5] = 1
	iMatrix[10] = 1
	iMatrix[15] = 1
	identityMatrix := mat64.NewDense(4, 4, iMatrix)
	identityMatrix.Add(identityMatrix, sine)
	identityMatrix.Add(identityMatrix, cosine)

	return identityMatrix
}

// NewScalingMatrix ...
func NewScalingMatrix(x, y, z float64) *mat64.Dense {
	matrix := make([]float64, 16)
	matrix[0] = x
	matrix[5] = y
	matrix[10] = z
	matrix[15] = 1
	return mat64.NewDense(4, 4, matrix)
}

// Transform ...
func Transform(matrix *mat64.Dense, vec primitives.Vec3) primitives.Vec3 {
	vector := make([]float64, 4)
	vector[0] = vec.X()
	vector[1] = vec.Y()
	vector[2] = vec.Z()
	vector[3] = 1
	vMatrix := mat64.NewDense(4, 1, vector)
	out := mat64.NewDense(4, 1, nil)
	out.Mul(matrix, vMatrix)
	return primitives.NewVec3(out.At(0, 0), out.At(1, 0), out.At(2, 0))
}

// TransformNormal ...
func TransformNormal(matrix *mat64.Dense, vec primitives.Vec3) primitives.Vec3 {
	vector := make([]float64, 4)
	vector[0] = vec.X()
	vector[1] = vec.Y()
	vector[2] = vec.Z()
	vMatrix := mat64.NewDense(4, 1, vector)
	out := mat64.NewDense(4, 1, nil)
	out.Mul(matrix.T(), vMatrix)
	return primitives.NewVec3(out.At(0, 0), out.At(1, 0), out.At(2, 0))
}

// TransformRay changes modifies the ray by the transformation
func TransformRay(matrix *mat64.Dense, ray *primitives.Ray) *primitives.Ray {
	tdirection := Transform(matrix, ray.Direction())
	torigin := Transform(matrix, ray.Origin())
	return primitives.NewRay(torigin, tdirection)
}

// Coalesce ...
func Coalesce(matrices []*mat64.Dense) *mat64.Dense {
	iMatrix := make([]float64, 16)
	iMatrix[0] = 1
	iMatrix[5] = 1
	iMatrix[10] = 1
	iMatrix[15] = 1
	result := mat64.NewDense(4, 4, iMatrix)
	for _, matrix := range matrices {
		result.Mul(result, matrix)
	}
	return result
}
