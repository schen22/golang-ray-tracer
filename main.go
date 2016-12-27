package main

import "fmt"

/** Linearly blend white and blue
depending on up/downess of the y-coordinate. */
func color(r Ray) Vec3 {
	unitDirection := r.direction.FindUnitVector()
	// Scale unit vector so that 0.0 < t < 1.0
	var t float64 = 0.5 * (unitDirection.y() + 1.0)
	// Form a linear interpolation between blue to white
	startVal := Vec3{1, 1, 1}.MultiplyNum(1.0 - t)
	endVal := Vec3{0.5, 0.7, 1.0}.MultiplyNum(t)
	blendedVal := startVal.AddVector(endVal)
	return blendedVal
}

func main() {
	nx := 200
	ny := 100

	lowerLeftCorner := Vec3{-2, -1, -1}
	horizontal := Vec3{4, 0, 0}
	vertical := Vec3{0, 2, 0}
	origin := Vec3{0, 0, 0}
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			horizontal = horizontal.MultiplyNum(u)
			vertical = vertical.MultiplyNum(v)
			direction := lowerLeftCorner.AddVector(horizontal).AddVector(vertical)
			ray := Ray{origin, direction}
			col := color(ray)
			ir := int(255.99 * col.x())
			ig := int(255.99 * col.y())
			ib := int(255.99 * col.z())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
