package utils

// PerlinNoise ...
type PerlinNoise struct {
	px, py, pz int
	random     float64
}

// func NewPerlinNoise(seed primitives.Vec3) PerlinNoise {
//   u := seed.X() - math.Floor(seed.X())
//   v := seed.Y() - math.Floor(seed.Y())
//   o := seed.Z() - math.Floor(seed.Z())
//   i := math.Floor(seed.X())
//   j := math.Floor(seed.Y())
//   k := math.Floor(seed.Z())
// }
