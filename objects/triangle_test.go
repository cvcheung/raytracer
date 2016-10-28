package objects

import (
	"fmt"
	"testing"
	"raytracer/primitives"
	"raytracer/materials"
	"raytracer/textures"
)

func TestTriangleHit(t *testing.T) {
	zero := primitives.NewVec3(0.0, 0.0, 0.0)
	v0 := primitives.NewVec3(0.0, 0.0, -2.0)
	v1 := primitives.NewVec3(0.0, 1.0, -2.0)
	v2 := primitives.NewVec3(1.0, 0.0, -2.0)
	l := materials.NewAmbientLight(textures.Red)
	mat := materials.NewBlinnphong(textures.Blue, textures.Blue, textures.Blue, textures.Blue, 16.0, l)
	var rset [100]*primitives.Ray
	index := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			rset[index] = primitives.NewRay(primitives.NewVec3(0.0, 0.0, 1.0), primitives.NewVec3( 0.0 + float64(i) / 10.0, 0.0 + float64(j) / 10.0, -3.0))
			index = index + 1
		}
	}

	tr := NewTriangle(v0, v1, v2, mat)
	hitRecord := materials.NewRecord(0.0, 0.0, 0.0, zero, zero, mat)
	var hitting [100]bool
	var rays [100]string
	for in := 0; in < 100; in++ {
		hitting[in] = tr.Hit(rset[in], 0.0000001, 100.0, hitRecord)
		rays[in] = rset[in].String()
	}
	fmt.Println(hitting)

}

func TestTriangleMiss(t *testing.T) {
	zero := primitives.NewVec3(0.0, 0.0, 0.0)
	v0 := primitives.NewVec3(0.0, 0.0, -2.0)
	v1 := primitives.NewVec3(0.0, 1.0, -2.0)
	v2 := primitives.NewVec3(1.0, 0.0, -2.0)
	l := materials.NewAmbientLight(textures.Red)
	mat := materials.NewBlinnphong(textures.Blue, textures.Blue, textures.Blue, textures.Blue, 16.0, l)
	r := primitives.NewRay(primitives.NewVec3(0.5, -1.0, 0.2), primitives.NewVec3(-0.2, 1.0, 0.1))

	fmt.Println(v0)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(r)
	tri := NewTriangle(v0, v1, v2, mat)
	hitRecord := materials.NewRecord(0.0, 0.0, 0.0, zero, zero, mat)
	ifhit := tri.Hit(r, 0.00001, 100.0, hitRecord)
	fmt.Println(ifhit)
	if ifhit {
		t.Error("Should have missed")
	}
	fmt.Println(hitRecord)
}