package materials

import "golang-ray-tracer/models"

// Material to define how rays interact with the surface.
type Material interface {

	// This function returns a boolean to determine if material scatter's the ray.
	// @param ray_in: ray interacting with surface
	// @param rec: contains pointer to the object's material, point vector and its normal.
	// @param attenuation: the amount of light reflection reduced to.
	// @param scattered:
	Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool
}
