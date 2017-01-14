package objects

import (
	"golang-ray-tracer/models"
	"math"
)

type Sphere struct {
	Center models.Vec3
	Radius float64
}

func (s Sphere) Hit(r models.Ray, t_min float64, t_max float64, rec *HitRecord) bool {
	oc := r.Origin.SubtractVector(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < t_max && temp > t_min {
			rec.T = temp
			rec.P = r.PointAtParam(rec.T)
			rec.Normal = rec.P.SubtractVector(s.Center).DivideNum(s.Radius)
			return true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < t_max && temp > t_min {
			rec.T = temp
			rec.P = r.PointAtParam(rec.T)
			rec.Normal = rec.P.SubtractVector(s.Center).DivideNum(s.Radius)
			return true
		}
	}
	return false
}
