package materials

import "golang-ray-tracer/models"

// Material that diffuses rays. Light is reflected at many angles.
// param Albedo - proportion of light reflected by the material.
type Lambertian struct {
	Albedo models.Vec3
}

func (l Lambertian) Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool {
	target := rec.P.AddVector(rec.Normal).AddVector(models.RandomInUnitSphere())
	*scattered = models.Ray{Origin: rec.P, Direction: target.SubtractVector(rec.P)}
	*attenuation = l.Albedo
	return true
}
