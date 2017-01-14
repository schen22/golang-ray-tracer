package models

type Ray struct {
	Origin, Direction Vec3
}

func NewRay(origin Vec3, direction Vec3) Ray {
	return Ray{origin, direction}
}

/* Compute the position of a point along a 3D space . */
func (r Ray) PointAtParam(t float64) Vec3 {
	var v Vec3 = r.Direction.MultiplyNum(t)
	return r.Origin.AddVector(v)
}
