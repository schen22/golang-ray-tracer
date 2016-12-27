package main

type Ray struct {
	origin    Vec3
	direction Vec3
}

/* Compute the position of a point along a 3D space . */
func (r Ray) PointAtParam(t float64) Vec3 {
	var v Vec3 = r.direction.MultiplyNum(t)
	return r.origin.AddVector(v)
}
