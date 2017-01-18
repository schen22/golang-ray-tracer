package materials

import (
	"fmt"
	"golang-ray-tracer/models"
)

type Lambertian struct {
	Albedo models.Vec3
}

func (l Lambertian) Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool {
	target := rec.P.AddVector(rec.Normal).AddVector(models.RandomInUnitSphere())
	*scattered = models.Ray{Origin: rec.P, Direction: target.SubtractVector(rec.P)}
	fmt.Printf("scattered = %v", *scattered)
	*attenuation = l.Albedo
	return true
}
