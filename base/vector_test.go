package base

import (
	"fmt"
	"math"
	"testing"
)

func TestVector(t *testing.T) {
	Lena := NewVec3(7.0, 17.0, 27.0)
	Kevin := NewVec3(6.0, 9.0, 69.0)
	fmt.Println(Lena)
	fmt.Println(Kevin)
	Adi := Kevin.Magnitude() - Lena.Magnitude()
	fmt.Println(Adi)
	Lena.Normalize()
	fmt.Println(Lena)
}

func TestAddition(t *testing.T) {
	Lena := NewVec3(7.0, 17.0, 27.0)
	Kevin := NewVec3(6.0, 9.0, 69.0)
	Adi := Lena.Subtract(Kevin)
	if Adi.x != 1 || Adi.y != 8 || Adi.z != -42 {
		t.Error("Failed Addition")
	}
}

func TestDotProduct(t *testing.T) {
	a := NewVec3(3, -3, 1)
	b := NewVec3(4, 9, 2)
	if a.Dot(b) != float64(12-27+2) {
		t.Error(a, b)
		t.Errorf("%f != %f", a.Dot(b), float64(12-27+2))
	}

}

func TestCrossProduct(t *testing.T) {
	a := NewVec3(3, -3, 1)
	b := NewVec3(4, 9, 2)
	axb := a.Cross(b)
	if axb.Magnitude() != 5*math.Sqrt(70) {
		t.Error(axb)
		t.Errorf("%f != %f\n", axb.Magnitude(), 5*math.Sqrt(70))
	}
}
