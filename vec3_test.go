package main

import (
	"testing"
)

func TestVec3Multiply(t *testing.T) {
	var vec Vec3
	vec = Vec3{1, 0.1, -2}
	expectedVec := Vec3{1, 0.1, -2}
	if vec != expectedVec {
		t.Error("Expected {1 0.1 -2}, got ", vec)
	}

	vec1 := vec.multiply(2)
	expectedVec1 := Vec3{2, 0.2, -4}
	if vec1 != expectedVec1 {
		t.Error("Expected {2 0.2 -4}, got ", vec1)
	}
}
