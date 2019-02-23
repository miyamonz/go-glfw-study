package main

type Vec3 [3]float32
type Vec4 [4]float32

func (v *Vec3) add(v2 Vec3) Vec3 {
	return Vec3{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
}
func (v *Vec3) sub(v2 Vec3) Vec3 {
	return Vec3{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
}
func (v *Vec3) mult(n float32) Vec3 {
	return Vec3{
		v[0] * n,
		v[1] * n,
		v[2] * n,
	}
}
func (v *Vec3) cross(v2 Vec3) Vec3 {
	return Vec3{
		v[1]*v2[2] - v[2]*v2[1],
		v[2]*v2[0] - v[0]*v2[2],
		v[0]*v2[1] - v[1]*v2[0],
	}
}

func (v *Vec3) length() float32 {
	return sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}
