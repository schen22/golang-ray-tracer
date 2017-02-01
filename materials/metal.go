package materials

import "golang-ray-tracer/models"

type Metal struct {
	Albedo models.Vec3
	Fuzz   float64
}

/** Rays are reflected as (v + 2B) for smooth metals. */
func Reflect(v models.Vec3, n models.Vec3) models.Vec3 {
	return v.SubtractVector(n.MultiplyNum(2.0 * v.Dot(n)))
}

func (m Metal) Scatter(ray_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool {
	reflected := Reflect(ray_in.Direction.FindUnitVector(), rec.Normal)
	// Put a max of 1 on the radius of the sphere to add fuzzier reflections
	// for larger spheres.
	if m.Fuzz > 1 {
		m.Fuzz = 1.0
	}
	// Randomize the reflected direction by using a small sphere to choose
	// a new endpoint for the ray.
	fuzzFactor := models.RandomInUnitSphere().MultiplyNum(m.Fuzz)
	*scattered = models.NewRay(rec.P, reflected.AddVector(fuzzFactor))
	*attenuation = m.Albedo
	return (scattered.Direction.Dot(rec.Normal) > 0)
}
