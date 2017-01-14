package objects

import "golang-ray-tracer/models"

type Camera struct {
	Origin, LowerLeftCorner models.Vec3
	Horizontal, Vertical    models.Vec3
}

func NewCamera(origin models.Vec3, llc models.Vec3, horiz models.Vec3, vert models.Vec3) Camera {
	return Camera{origin, llc, horiz, vert}
}

func (c Camera) GetRay(u float64, v float64) models.Ray {
	horizontalOffset := c.Horizontal.MultiplyNum(u)
	verticalOffset := c.Vertical.MultiplyNum(v).SubtractVector(c.Origin)
	return models.NewRay(c.Origin, c.LowerLeftCorner.AddVector(horizontalOffset).AddVector(verticalOffset))
}
