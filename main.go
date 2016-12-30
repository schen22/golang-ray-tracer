package main

import (
	"fmt"
	"golang-ray-tracer/models"
	"math"
)

// Calculate when any pixel hits the sphere (represented by the
// center point of a vector).
func hitSphere(center models.Vec3, radius float64, ray models.Ray) float64 {
	oc := ray.Origin.SubtractVector(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}

/** Linearly blend white and blue
depending on up/downess of the y-coordinate. */
func color(r models.Ray) models.Vec3 {
	// If any pixel hits the sphere placed at -1 on the z-axis,
	// color the pixel by mapping surface normal to its RGB value.
	sphere := models.Vector(0, 0, -1)
	t := hitSphere(sphere, 0.5, r)
	if t > 0.0 {
		// Make surface normals unit vectors and normalize.
		N := r.PointAtParam(t).SubtractVector(sphere).FindUnitVector()
		return models.Vector(N.X()+1, N.Y()+1, N.Z()+1).MultiplyNum(0.5)
	}
	unitDirection := r.Direction.FindUnitVector()
	// Scale unit vector so that 0.0 < t < 1.0
	t = 0.5 * (unitDirection.Y() + 1.0)
	// Form a linear interpolation between blue to white
	startVal := models.Vector(1.0, 1.0, 1.0).MultiplyNum(1.0 - t)
	endVal := models.Vector(0.5, 0.7, 1.0).MultiplyNum(t)
	blendedVal := startVal.AddVector(endVal)
	return blendedVal
}

func main() {
	nx := 200
	ny := 100

	lowerLeftCorner := models.Vector(-2, -1, -1)
	horizontal := models.Vector(4, 0, 0)
	vertical := models.Vector(0, 2, 0)
	origin := models.Vector(0, 0, 0)
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			scaleHoriz := horizontal.MultiplyNum(u)
			scaleVert := vertical.MultiplyNum(v)
			direction := lowerLeftCorner.AddVector(scaleVert).AddVector(scaleHoriz)
			ray := models.Ray{origin, direction}
			col := color(ray)
			ir := int(255.99 * col.E0)
			ig := int(255.99 * col.E1)
			ib := int(255.99 * col.E2)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
