package base

// Object defines base class for 3D objects.
type Object interface {
	Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool
}

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

// NewEmptyObjectList ...
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
func (o *ObjectList) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	hit := false
	tempRec := HitRecord{}
	closestSoFar := tMax
	for _, v := range o.objects {
		if v.Hit(r, tMin, closestSoFar, &tempRec) {
			hit = true
			closestSoFar = tempRec.t
			rec.CopyRecord(&tempRec)
		}
	}

	return hit
}

// HitRecord is a simple record to that records the information regarding where
// the ray hit.
type HitRecord struct {
	t         float64
	p, normal Vec3
	mat       Material
}

// T returns the t value that caused the ray to intersect the object.
func (rec *HitRecord) T() float64 {
	return rec.t
}

// Point returns the the point at which the ray intersected the object.
func (rec *HitRecord) Point() Vec3 {
	return rec.p
}

// Normal returns the normal at which the ray intersected the object.
func (rec *HitRecord) Normal() Vec3 {
	return rec.normal
}

// Material returns the pointer to the material struct that defines the type of
// material that the ray hit.
func (rec *HitRecord) Material() Material {
	return rec.mat
}

// CopyRecord update the current record with the fields another record.
func (rec *HitRecord) CopyRecord(rec2 *HitRecord) {
	rec.t = rec2.t
	rec.p = rec2.p
	rec.normal = rec2.normal
	rec.mat = rec2.mat
}
