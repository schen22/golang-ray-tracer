package models

import "testing"

func TestVec3MultiplyNum(t *testing.T) {
	vec := Vector(1, 0.1, -2)
	expectedVec := Vector(1, 0.1, -2)
	if vec != expectedVec {
		t.Error("Expected ", expectedVec, "got ", vec)
	}

	vec1 := vec.MultiplyNum(2)
	expectedVec1 := Vector(2, 0.2, -4)
	if vec1 != expectedVec1 {
		t.Error("Expected ", expectedVec1, ", got ", vec1)
	}
}

func TestVec3DivideNum(t *testing.T) {
	vec := Vector(1, 2, 3)
	vec2 := vec.DivideNum(2)
	expected := Vector(0.5, 1, 1.5)
	// why does vec2 != vec?
	if vec2 != expected {
		t.Error("Expected: ", expected, "Got: ", vec2)
	}
}

func TestVec3FindUnitVector(t *testing.T) {
	vec1 := Vector(1, 2, 2)
	u := vec1.FindUnitVector()
	expectedVec1 := Vector(float64(1)/float64(3), float64(2)/float64(3), float64(2)/float64(3))
	// fmt.Println("unit vector = ", expectedVec1)
	if u != expectedVec1 {
		t.Error("Expected", expectedVec1, "got", u)
	}
}
