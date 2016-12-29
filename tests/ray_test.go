package main

import (
	"fmt"
	"testing"
)

func TestRayPointAtParam(t *testing.T) {
	origin := Vec3{0, 0, 0}
	direction := Vec3{1, 1, 1}
	ray := Ray{origin, direction}

	if ray.origin != origin {
		t.Error("Expected: ", origin, "Got: ", ray.origin)
	}

	if ray.direction != direction {
		t.Error("Expected: ", direction, "Got: ", ray.direction)
	}

	expected := Vec3{4, 4, 4}
	point := ray.PointAtParam(float64(4))
	if point != expected {
		t.Error("Expected ", expected, "Got: ", point)
	}

	expected2 := Vec3{2.2, 2.2, 2.2}
	point2 := ray.PointAtParam(2.2)
	if point2 != expected2 {
		t.Error("Expected ", expected2, "Got: ", point2)
	}
	// direction = direction.MultiplyNum(4.2)
	// fmt.Printf("direction = ", direction)
	//
	// direction = direction.AddVector(expected2)
	// fmt.Printf("origin = ", origin)
	// fmt.Printf("direction 2 = ", direction)
	// ray2 := Ray{origin, direction}
	// fmt.Printf("ray = ", ray2)
	vecEx := Vec3{-2, 0.98, 7}
	fmt.Printf("unit vector = ", vecEx.FindUnitVector())
}
