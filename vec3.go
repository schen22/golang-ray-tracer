package main

import "math"

/// Structs are typed collections of fields useful for grouping
/// data together to form records.

type Vec3 struct {
	e0 float64
	e1 float64
	e2 float64
}

//// vec3(float e0, float e1, float e2) { e[0] = e0; e[1] = e1; e[2] = e2;}
func (v *Vec3) x() float64 {
	return v.e0
}
func (v *Vec3) y() float64 {
	return v.e1
}
func (v *Vec3) z() float64 {
	return v.e2
}

func (v *Vec3) AddNum(num float64) Vec3 {
	return Vec3{v.e0 + num, v.e1 + num, v.e2 + num}
}
func (v *Vec3) SubtractNum(num float64) Vec3 {
	return Vec3{v.e0 - num, v.e1 - num, v.e2 - num}
}

func (v *Vec3) MultiplyNum(num float64) Vec3 {
	return Vec3{v.e0 * num, v.e1 * num, v.e2 * num}
}
func (v *Vec3) DivideNum(num float64) Vec3 {
	return Vec3{v.e0 / num, v.e1 / num, v.e2 / num}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2)
}
func (v *Vec3) SquaredLength() float64 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}
func (v *Vec3) FindUnitVector() Vec3 {
	l := v.Length()
	u := v.DivideNum(l)
	return u
}
func (v0 *Vec3) AddVector(v1 *Vec3) Vec3 {
	return Vec3{v0.e0 + v1.e0, v0.e1 + v1.e1, v0.e2 + v1.e2}
}
func (v0 *Vec3) SubtractVector(v1 *Vec3) Vec3 {
	return Vec3{v0.e0 - v1.e0, v0.e1 - v1.e1, v0.e2 - v1.e2}
}

func (v0 *Vec3) MultiplyVector(v1 *Vec3) Vec3 {
	return Vec3{v0.e0 * v1.e0, v0.e1 * v1.e1, v0.e2 * v1.e2}
}
func (v0 *Vec3) DivideVector(v1 *Vec3) Vec3 {
	return Vec3{v0.e0 / v1.e0, v0.e1 / v1.e1, v0.e2 / v1.e2}
}
func (v0 *Vec3) DotProduct(v1 *Vec3) float64 {
	return v0.e0*v1.e0 + v0.e1*v1.e1 + v0.e2*v1.e2
}
func (v0 *Vec3) Cross(v1 *Vec3) Vec3 {
	return Vec3{v0.e1*v1.e2 - v0.e2*v1.e1,
		-(v0.e0*v1.e2 - v0.e2*v1.e0),
		v0.e0*v1.e1 - v0.e1*v1.e0}
}

// func (v *Vec3)

/// inline const vec3& operator+ () const { return *this; }
/// inline vec3 operator-() const { return vec3(-e[0], -e[1], -e[2]); }
/// inline float operator[](int i) const { return e[i];};
