package objects

import (
	"log"
	"sort"
)

// Taken from the Golang sort example.

// By is the type of a "less" function that defines the ordering of its object arguments.
type By func(a, b Object) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(objects []Object) {
	os := &objectSorter{
		objects: objects,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(os)
}

// objectSorter joins a By function and a slice of Objects to be sorted.
type objectSorter struct {
	objects []Object
	by      func(a, b Object) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *objectSorter) Len() int {
	return len(s.objects)
}

// Swap is part of sort.Interface.
func (s *objectSorter) Swap(i, j int) {
	s.objects[i], s.objects[j] = s.objects[j], s.objects[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *objectSorter) Less(i, j int) bool {
	return s.by(s.objects[i], s.objects[j])
}

// BoxCompareX ...
func BoxCompareX(a, b Object) bool {
	leftHit, leftBox := a.BoundingBox(0, 0)
	rightHit, rightBox := b.BoundingBox(0, 0)
	if !leftHit || !rightHit {
		log.Println("no bounding box in BVHNode constructor")
	}
	return leftBox.min.X() < rightBox.min.X()
}

// BoxCompareY ...
func BoxCompareY(a, b Object) bool {
	leftHit, leftBox := a.BoundingBox(0, 0)
	rightHit, rightBox := b.BoundingBox(0, 0)
	if !leftHit || !rightHit {
		log.Println("no bounding box in BVHNode constructor")
	}
	return leftBox.min.Y() < rightBox.min.Y()
}

// BoxCompareZ ...
func BoxCompareZ(a, b Object) bool {
	leftHit, leftBox := a.BoundingBox(0, 0)
	rightHit, rightBox := b.BoundingBox(0, 0)
	if !leftHit || !rightHit {
		log.Println("no bounding box in BVHNode constructor")
	}
	return leftBox.min.Z() < rightBox.min.Z()
}
