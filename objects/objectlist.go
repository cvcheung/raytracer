package objects

import (
	"raytracer/materials"
	"raytracer/primitives"
)

// ObjectList is a container for all the objects that our ray can hit in the
// scene.
type ObjectList struct {
	objects []Object
}

// NewObjectList constructs an ObjectList from the objects given.
func NewObjectList(length int, objs ...Object) *ObjectList {
	objList := make([]Object, length)
	for i, v := range objs {
		objList[i] = v
	}
	return &ObjectList{objList}
}

// NewEmptyObjectList returns a new empty ObjectList for efficient adding.
func NewEmptyObjectList(length int) *ObjectList {
	objList := make([]Object, 0, length)
	return &ObjectList{objList}
}

// Add an object to the object list.
func (o *ObjectList) Add(obj Object) {
	o.objects = append(o.objects, obj)
}

// Hit for an ObjectList iterates through the objects inside the list and puts
// into the record the first object that would be hit.
func (o *ObjectList) Hit(r *primitives.Ray, tMin, tMax float64, rec *materials.HitRecord) bool {
	hit := false
	tempRec := materials.HitRecord{}
	closestSoFar := tMax
	for _, v := range o.objects {
		if v.Hit(r, tMin, closestSoFar, &tempRec) {
			hit = true
			closestSoFar = tempRec.T()
			rec.CopyRecord(&tempRec)
		}
	}
	return hit
}
