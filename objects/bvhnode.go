package objects

import (
	"raytracer/materials"
	"raytracer/primitives"
)

// BVHNode is a container that defines the search space for our rays using AABB
// to reduce to computation time per pixel from linear to the number of objects
// to sublinear.
type BVHNode struct {
	box         *AABB
	left, right *BVHNode
}

// bvh_node::bvh_node(hitable **l, int n, float time0, float time1) {
//     int axis = int(3*drand48());
//     if (axis == 0)
//        qsort(l, n, sizeof(hitable *), box_x_compare);
//     else if (axis == 1)
//        qsort(l, n, sizeof(hitable *), box_y_compare);
//     else
//        qsort(l, n, sizeof(hitable *), box_z_compare);
//     if (n == 1) {
//         left = right = l[0];
//     }
//     else if (n == 2) {
//         left = l[0];
//         right = l[1];
//     }
//     else {
//         left = new bvh_node(l, n/2, time0, time1);
//         right = new bvh_node(l + n/2, n - n/2, time0, time1);
//     }
//     aabb box_left, box_right;
//     if(!left->bounding_box(time0,time1, box_left) || !right->bounding_box(time0,time1, box_right))
//         std::cerr << "no bounding box in bvh_node constructor\n";
//     box = surrounding_box(box_left, box_right);
// }

// func NewBVHNode(hitable *Object, n int, t0, t1 float64) *BVHNode {
//   return &BVHNode{box, left, right}
// }

// NewEmptyBVHNode returns a BVHNode with its fields undeclared.
func NewEmptyBVHNode() *BVHNode {
	return &BVHNode{}
}

// Hit ...
func (n *BVHNode) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	if n.box.Hit(r, tMin, tMax, rec) {
		var leftRec, rightRec *materials.HitRecord
		leftHit := n.left.Hit(r, tMin, tMax, leftRec)
		rightHit := n.right.Hit(r, tMin, tMax, rightRec)
		if leftHit && rightHit {
			if leftRec.T() < rightRec.T() {
				rec.CopyRecord(leftRec)
			} else {
				rec.CopyRecord(rightRec)
			}
			return true
		} else if leftHit {
			rec.CopyRecord(leftRec)
			return true
		} else if rightHit {
			rec.CopyRecord(rightRec)
			return true
		}
		return false
	}
	return false
}
