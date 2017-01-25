package models

import (
	"fmt"
	"testing"
)

func TestRayPointAtParam(t *testing.T) {
	origin := Vector(0, 0, 0)
	direction := Vector(1, 1, 1)
	ray := Ray{Origin: origin, Direction: direction}

	if ray.Origin != origin {
		t.Error("Expected: ", origin, "Got: ", ray.Origin)
	}

	if ray.Direction != direction {
		t.Error("Expected: ", direction, "Got: ", ray.Direction)
	}

	expected := Vector(4, 4, 4)
	point := ray.PointAtParam(float64(4))
	if point != expected {
		t.Error("Expected ", expected, "Got: ", point)
	}

	expected2 := Vector(2.2, 2.2, 2.2)
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
	vecEx := Vector(-2, 0.98, 7)
	fmt.Printf("unit vector = %v", vecEx.FindUnitVector())
}
