package materials

import "golang-ray-tracer/models"

type Metal struct {
	Albedo models.Vec3
}

func Reflect(v models.Vec3, n models.Vec3) models.Vec3 {
	return v.SubtractVector(n.MultiplyNum(2.0 * v.Dot(n)))
}

func (m Metal) Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool {
	reflected := Reflect(ray_in.Direction.FindUnitVector(), rec.Normal)
	*scattered = models.NewRay(rec.P, reflected)
	*attenuation = m.Albedo
	return (scattered.Direction.Dot(rec.Normal) > 0)
}
