package main

import (
	"fmt"
	"golang-ray-tracer/models"
	"golang-ray-tracer/objects"
	"math"
)

/** Linearly blend white and blue
depending on up/downess of the y-coordinate. */
func color(r models.Ray, world objects.Hitable) models.Vec3 {
	var rec objects.HitRecord

	if world.Hit(r, 0.0, math.MaxFloat64, &rec) {
		// fmt.Printf("rec normal object = %v", rec.Normal.X())

		return models.Vector(rec.Normal.X()+1, rec.Normal.Y()+1, rec.Normal.Z()+1).MultiplyNum(0.5)
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
	nx := 200
	ny := 100

	lowerLeftCorner := models.Vector(-2, -1, -1)
	horizontal := models.Vector(4, 0, 0)
	vertical := models.Vector(0, 2, 0)
	origin := models.Vector(0, 0, 0)
	list := make([]objects.Hitable, 2)
	list[0] = objects.Sphere{Center: models.Vector(0, 0, -1), Radius: 0.5}
	list[1] = objects.Sphere{Center: models.Vector(0, -100.5, -1), Radius: 100}
	world := objects.HitableList{List: list, ListSize: 2}
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			scaleHoriz := horizontal.MultiplyNum(u)
			scaleVert := vertical.MultiplyNum(v)
			direction := lowerLeftCorner.AddVector(scaleVert).AddVector(scaleHoriz)
			ray := models.Ray{Origin: origin, Direction: direction}

			// p := ray.PointAtParam(2.0)
			col := color(ray, world)
			ir := int(255.99 * col.E0)
			ig := int(255.99 * col.E1)
			ib := int(255.99 * col.E2)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
