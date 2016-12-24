package main

/// Structs are typed collections of fields useful for grouping
/// data together to form records.

type Vec3 struct {
	e0 float64
	e1 float64
	e2 float64
}

func (v *Vec3) x() float64 {
	return v.e0
}
func (v *Vec3) y() float64 {
	return v.e1
}
func (v *Vec3) z() float64 {
	return v.e2
}
func (v *Vec3) r() float64 {
	return v.e0
}
func (v *Vec3) g() float64 {
	return v.e1
}
func (v *Vec3) b() float64 {
	return v.e2
}

func (v *Vec3) multiply(num float64) Vec3 {
	return Vec3{v.e0 * num, v.e1 * num, v.e2 * num}
}

/// inline const vec3& operator+ () const { return *this; }
/// inline vec3 operator-() const { return vec3(-e[0], -e[1], -e[2]); }
/// inline float operator[](int i) const { return e[i];};
