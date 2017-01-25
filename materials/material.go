package materials

import "golang-ray-tracer/models"

// Material to define how rays interact with the surface.
type Material interface {
	Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool
}
