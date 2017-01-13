package objects

import "golang-ray-tracer/models"

// t = time to allow motion blur.
// p = point vector, normal = the normal of that vector
type HitRecord struct {
	T         float64
	P, Normal models.Vec3
}

type Hitable interface {
	// Determine if ray object will be hit within the interval
	// tmin < t < tmax.
	Hit(r models.Ray, t_min float64, t_max float64, rec *HitRecord) bool
}
