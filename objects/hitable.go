package models

type hitRecord struct {
	t      float64
	p      Vec3
	normal Vec3
}

type hitable interface {
  func hit(r Ray, t_min float64, t_max float64, rec hitRecord) bool {

  }
}
