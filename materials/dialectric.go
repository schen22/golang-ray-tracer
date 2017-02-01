package materials

import (
	"golang-ray-tracer/models"
	"math"
	"math/rand"
)

type Dialectric struct {
	Ref_idx float64
}

func Refract(v *models.Vec3, n *models.Vec3, ni_over_nt float64, refracted *models.Vec3) bool {
	uv := v.FindUnitVector()
	dt := uv.Dot(*n)
	discriminant := 1.0 - ni_over_nt*ni_over_nt*(1-dt*dt)
	if discriminant > 0 {
		refract_uv := uv.SubtractVector(n.MultiplyNum(dt)).MultiplyNum(ni_over_nt)
		refract_n := n.MultiplyNum(math.Sqrt(discriminant))
		*refracted = refract_uv.SubtractVector(refract_n)
		return true
	} else {
		return false
	}
}

func (d Dialectric) Scatter(r_in *models.Ray, rec *HitRecord, attenuation *models.Vec3, scattered *models.Ray) bool {
	var outward_normal models.Vec3
	var ni_over_nt float64
	var refracted models.Vec3
	var reflect_prob float64
	var cosine float64
	reflected := Reflect(r_in.Direction, rec.Normal)
	*attenuation = models.Vector(1.0, 1.0, 1.0)
	if r_in.Direction.Dot(rec.Normal) > 0 {
		outward_normal = rec.Normal.MultiplyNum(-1)
		ni_over_nt = d.Ref_idx
		cosine = d.Ref_idx * r_in.Direction.Dot(rec.Normal) / r_in.Direction.Length()
	} else {
		outward_normal = rec.Normal
		ni_over_nt = 1.0 / d.Ref_idx
		cosine = -(r_in.Direction.Dot(rec.Normal)) / r_in.Direction.Length()
	}
	if Refract(&r_in.Direction, &outward_normal, ni_over_nt, &refracted) {
		*scattered = models.NewRay(rec.P, refracted)
		reflect_prob = Schlick(cosine, d.Ref_idx)
	} else {
		*scattered = models.NewRay(rec.P, reflected)
		reflect_prob = 1.0
	}
	if rand.Float64() < reflect_prob {
		*scattered = models.NewRay(rec.P, reflected)
	} else {
		*scattered = models.NewRay(rec.P, refracted)
	}
	return true
}

/** Reflection varies at angle of ray hitting the surface. Polynomial
  approximation by Shclick. */
func Schlick(cosine float64, ref_idx float64) float64 {
	r0 := (1 - ref_idx) / (1 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
