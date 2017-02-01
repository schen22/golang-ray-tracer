package main

import (
	"fmt"
	"golang-ray-tracer/materials"
	"golang-ray-tracer/models"
	"golang-ray-tracer/objects"
	"math"
	"math/rand"
)

/** Linearly blend white and blue
depending on up/downess of the y-coordinate. */
func color(r models.Ray, world materials.Hitable, depth int) models.Vec3 {
	var rec materials.HitRecord
	if world.Hit(r, 0.001, math.MaxFloat64, &rec) {
		var scattered models.Ray
		var attenuation models.Vec3

		if depth < 50 && rec.MaterialPtr.Scatter(&r, &rec, &attenuation, &scattered) {
			return attenuation.MultiplyVector(color(scattered, world, depth+1))
		} else {
			return models.Vector(0, 0, 0)
		}
	} else {
		unitDirection := r.Direction.FindUnitVector()
		// Scale unit vector so that 0.0 < t < 1.0
		t := 0.5 * (unitDirection.Y() + 1.0)
		// Form a linear interpolation between blue to white
		startVal := models.Vector(1.0, 1.0, 1.0).MultiplyNum(1.0 - t)
		endVal := models.Vector(0.5, 0.7, 1.0).MultiplyNum(t)
		blendedVal := startVal.AddVector(endVal)
		return blendedVal
	}
}

func main() {
	nx := 600
	ny := 300
	ns := 100

	lowerLeftCorner := models.Vector(-2, -1, -1)
	horizontal := models.Vector(4, 0, 0)
	vertical := models.Vector(0, 2, 0)
	origin := models.Vector(0, 0, 0)
	list := make([]materials.Hitable, 5)
	list[0] = objects.Sphere{Center: models.Vector(0, 0, -1), Radius: 0.5, Material: materials.Lambertian{Albedo: models.Vector(0.1, 0.2, 0.5)}}
	list[1] = objects.Sphere{Center: models.Vector(0, -100.5, -1), Radius: 100, Material: materials.Lambertian{Albedo: models.Vector(0.8, 0.8, 0.0)}}
	list[2] = objects.Sphere{Center: models.Vector(1, 0, -1), Radius: 0.5, Material: materials.Metal{Albedo: models.Vector(0.8, 0.6, 0.2)}}
	// list[3] = objects.Sphere{Center: models.Vector(-1, 0, -1), Radius: 0.5, Material: materials.Metal{Albedo: models.Vector(0.8, 0.8, 0.8)}}
	list[3] = objects.Sphere{Center: models.Vector(-1, 0, -1), Radius: 0.5, Material: materials.Dialectric{Ref_idx: 1.5}}

	list[4] = objects.Sphere{Center: models.Vector(-1, 0, -1), Radius: -0.45, Material: materials.Dialectric{Ref_idx: 1.5}}

	world := materials.HitableList{List: list, ListSize: 5}
	cam := objects.NewCamera(origin, lowerLeftCorner, horizontal, vertical)
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// Blend pixels
			col := models.Vector(0, 0, 0)
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				// p := r.PointAtParam(2.0)
				col = col.AddVector(color(r, world, 0))
			}
			col = col.DivideNum(float64(ns))
			col = models.Vector(math.Sqrt(col.E0), math.Sqrt(col.E1), math.Sqrt(col.E2))
			ir := int(255.99 * col.E0)
			ig := int(255.99 * col.E1)
			ib := int(255.99 * col.E2)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
