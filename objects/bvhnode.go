package objects

import (
	"math"
	"raytracer/materials"
	"raytracer/primitives"
)

// BVHNode is a container that defines the search space for our rays using AABB
// to reduce to computation time per pixel from linear to the number of objects
// to logarithmic.
type BVHNode struct {
	box         *AABB
	left, right Object
}

// NewBVHNode recursively constructs a BVH given a list of objects.
func NewBVHNode(hitable []Object, n int, t0, t1 float64) *BVHNode {
	boxes := make([]*AABB, n)
	leftArea := make([]float64, n)
	rightArea := make([]float64, n)
	_, mainBox := hitable[0].BoundingBox(t0, t1)
	for i := 1; i < n; i++ {
		_, tempBox := hitable[i].BoundingBox(t0, t1)
		mainBox = SurroundingBox(mainBox, tempBox)
	}
	axis := mainBox.LongestAxis()
	if axis == 0 {
		By(BoxCompareX).Sort(hitable)
	} else if axis == 1 {
		By(BoxCompareY).Sort(hitable)
	} else {
		By(BoxCompareZ).Sort(hitable)
	}

	for i, v := range hitable {
		_, boxes[i] = v.BoundingBox(t0, t1)
	}

	leftBox := boxes[0]
	leftArea[0] = leftBox.Area()
	for i := 1; i < n-1; i++ {
		leftBox = SurroundingBox(leftBox, boxes[i])
		leftArea[i] = leftBox.Area()
	}

	rightBox := boxes[n-1]
	rightArea[0] = rightBox.Area()
	for i := n - 2; i > 0; i-- {
		rightBox = SurroundingBox(rightBox, boxes[i])
		rightArea[i] = rightBox.Area()
	}

	minSAH := math.MaxFloat64
	var minSAHidx int
	for i := 0; i < n-1; i++ {
		SAH := float64(i)*leftArea[i] + (float64(n)-float64(i)-1)*rightArea[i+1]
		if SAH < minSAH {
			minSAH = SAH
			minSAHidx = i
		}
	}

	var left, right Object
	if minSAHidx == 0 {
		left = hitable[0]
	} else {
		left = NewBVHNode(hitable[:minSAHidx+1], minSAHidx+1, t0, t1)
	}
	if minSAHidx == (n - 2) {
		right = hitable[minSAHidx+1]
	} else {
		right = NewBVHNode(hitable[minSAHidx+1:], n-minSAHidx-1, t0, t1)
	}
	box := mainBox
	return &BVHNode{box, left, right}

	// if axis == 0 {
	// 	By(BoxCompareX).Sort(hitable)
	// } else if axis == 1 {
	// 	By(BoxCompareY).Sort(hitable)
	// } else {
	// 	By(BoxCompareZ).Sort(hitable)
	// }
	//
	// var left, right Object
	// if n == 1 {
	// 	left = hitable[0]
	// 	right = hitable[0]
	// } else if n == 2 {
	// 	left = hitable[0]
	// 	right = hitable[1]
	// } else {
	// 	left = NewBVHNode(hitable[:n/2], n/2, t0, t1)
	// 	right = NewBVHNode(hitable[n/2:], n-n/2, t0, t1)
	// }
	// leftHit, leftBox := left.BoundingBox(t0, t1)
	// rightHit, rightBox := right.BoundingBox(t0, t1)
	// if !leftHit || !rightHit {
	// 	log.Println("no bounding box in BVHNode constructor")
	// }
	// box := SurroundingBox(leftBox, rightBox)
	// return &BVHNode{box, left, right}
}

// NewEmptyBVHNode returns a BVHNode with its fields undeclared.
func NewEmptyBVHNode() *BVHNode {
	return &BVHNode{}
}

// Hit ...
func (n *BVHNode) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	if n.box.Hit(r, tMin, tMax, rec) {
		leftRec := materials.HitRecord{}
		rightRec := materials.HitRecord{}
		leftHit := n.left.Hit(r, tMin, tMax, &leftRec)
		rightHit := n.right.Hit(r, tMin, tMax, &rightRec)
		if leftHit && rightHit {
			if leftRec.T() < rightRec.T() {
				rec.CopyRecord(&leftRec)
			} else {
				rec.CopyRecord(&rightRec)
			}
			return true
		} else if leftHit {
			rec.CopyRecord(&leftRec)
			return true
		} else if rightHit {
			rec.CopyRecord(&rightRec)
			return true
		}
		return false
	}
	return false
}

// BoundingBox ...
func (n *BVHNode) BoundingBox(t0, t1 float64) (bool, *AABB) {
	return true, n.box
}
