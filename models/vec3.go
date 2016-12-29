package models

import "math"

/// Structs are typed collections of fields useful for grouping
/// data together to form records.

type Vec3 struct {
	E0, E1, E2 float64
}

func Vector(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

//// vec3(float X(), float Y(), float Z()) { e[0] = X(); e[1] = Y(); e[2] = Z();}
func (v Vec3) X() float64 {
	return v.E0
}
func (v Vec3) Y() float64 {
	return v.E1
}
func (v Vec3) Z() float64 {
	return v.E2
}

func (v Vec3) AddNum(num float64) Vec3 {
	return Vec3{v.X() + num, v.Y() + num, v.Z() + num}
}
func (v Vec3) SubtractNum(num float64) Vec3 {
	return Vec3{v.X() - num, v.Y() - num, v.Z() - num}
}

func (v Vec3) MultiplyNum(num float64) Vec3 {
	return Vec3{v.X() * num, v.Y() * num, v.Z() * num}
}
func (v Vec3) DivideNum(num float64) Vec3 {
	return Vec3{v.X() / num, v.Y() / num, v.Z() / num}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
}
func (v Vec3) SquaredLength() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}
func (v Vec3) FindUnitVector() Vec3 {
	l := v.Length()
	u := v.DivideNum(l)
	return u
}
func (v0 Vec3) AddVector(v1 Vec3) Vec3 {
	return Vec3{v0.X() + v1.X(), v0.Y() + v1.Y(), v0.Z() + v1.Z()}
}
func (v0 Vec3) SubtractVector(v1 Vec3) Vec3 {
	return Vec3{v0.X() - v1.X(), v0.Y() - v1.Y(), v0.Z() - v1.Z()}
}

func (v0 Vec3) MultiplyVector(v1 Vec3) Vec3 {
	return Vec3{v0.X() * v1.X(), v0.Y() * v1.Y(), v0.Z() * v1.Z()}
}
func (v0 Vec3) DivideVector(v1 Vec3) Vec3 {
	return Vec3{v0.X() / v1.X(), v0.Y() / v1.Y(), v0.Z() / v1.Z()}
}
func (v0 Vec3) Dot(v1 Vec3) float64 {
	return v0.X()*v1.X() + v0.Y()*v1.Y() + v0.Z()*v1.Z()
}
func (v0 Vec3) Cross(v1 Vec3) Vec3 {
	return Vec3{v0.Y()*v1.Z() - v0.Z()*v1.Y(),
		-(v0.X()*v1.Z() - v0.Z()*v1.X()),
		v0.X()*v1.Y() - v0.Y()*v1.X()}
}

// func (v Vec3)

/// inline const vec3& operator+ () const { return *this; }
/// inline vec3 operator-() const { return vec3(-e[0], -e[1], -e[2]); }
/// inline float operator[](int i) const { return e[i];};
