package objects

import "golang-ray-tracer/models"

type HitableList struct {
	List     []Hitable
	ListSize int
}

func (hitableList HitableList) Hit(r models.Ray, t_min float64, t_max float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := t_max
	for i := 0; i < hitableList.ListSize; i++ {
		if hitableList.List[i].Hit(r, t_min, closestSoFar, &tempRec) {
			// fmt.Printf("tempRec = %v", tempRec)

			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}
	return hitAnything
}
