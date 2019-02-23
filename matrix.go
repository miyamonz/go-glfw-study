package main

import (
	"fmt"
)

type Mat3 [9]float32

type Matrix struct {
	matrix [16]float32
}

func (m *Matrix) data() *[16]float32 {
	return &m.matrix
}
func (m *Matrix) String() string {
	a := m.matrix
	var ret = ""
	ret += fmt.Sprintf("%f, %f, %f, %f,\n", a[0], a[1], a[2], a[3])
	ret += fmt.Sprintf("%f, %f, %f, %f,\n", a[4], a[5], a[6], a[7])
	ret += fmt.Sprintf("%f, %f, %f, %f,\n", a[8], a[9], a[10], a[11])
	ret += fmt.Sprintf("%f, %f, %f, %f\n", a[12], a[13], a[14], a[15])
	return ret
}
func (m *Matrix) loadIdentity() {

	for i := range m.matrix {
		m.matrix[i] = 0
	}

	m.matrix[0] = 1
	m.matrix[5] = 1
	m.matrix[10] = 1
	m.matrix[15] = 1
}

func identity() Matrix {
	m := Matrix{}
	m.loadIdentity()
	return m
}

func (m1 Matrix) getNormalMat() Mat3 {
	m := m1.data()
	r := Mat3{}
	r[0] = m[5]*m[10] - m[6]*m[9]
	r[1] = m[6]*m[8] - m[4]*m[10]
	r[2] = m[4]*m[9] - m[5]*m[8]
	r[3] = m[9]*m[2] - m[10]*m[1]
	r[4] = m[10]*m[0] - m[8]*m[2]
	r[5] = m[8]*m[1] - m[9]*m[0]
	r[6] = m[1]*m[6] - m[2]*m[5]
	r[7] = m[2]*m[4] - m[0]*m[6]
	r[8] = m[0]*m[5] - m[1]*m[4]
	return r
}

func (m1 Matrix) mult(m2 Matrix) Matrix {
	t := Matrix{}

	for i := range m1.matrix {
		j := i & 3
		k := i &^ 3

		t.matrix[i] =
			m1.matrix[0+j]*m2.matrix[k+0] +
				m1.matrix[4+j]*m2.matrix[k+1] +
				m1.matrix[8+j]*m2.matrix[k+2] +
				m1.matrix[12+j]*m2.matrix[k+3]
	}

	return t
}

func translate(v Vec3) Matrix {
	t := Matrix{}
	t.loadIdentity()
	t.matrix[12] = v[0]
	t.matrix[13] = v[1]
	t.matrix[14] = v[2]
	return t
}

func scale(v Vec3) Matrix {
	t := Matrix{}
	t.loadIdentity()
	t.matrix[0] = v[0]
	t.matrix[5] = v[1]
	t.matrix[10] = v[2]
	return t
}

func rotate(a float32, v Vec3) Matrix {
	x, y, z := v[0], v[1], v[2]
	d := sqrt(x*x + y*y + z*z)
	l := (x / d)
	m := (y / d)
	n := (z / d)

	l2 := (l * l)
	m2 := (m * m)
	n2 := (n * n)

	lm := (l * m)
	mn := (m * n)
	nl := (n * l)

	c := cos(a)
	c1 := (1.0 - c)
	s := sin(a)

	t := Matrix{}
	t.loadIdentity()
	t.matrix[0] = (1.0-l2)*c + l2
	t.matrix[1] = lm*c1 + n*s
	t.matrix[2] = nl*c1 - m*s
	t.matrix[4] = lm*c1 - n*s
	t.matrix[5] = (1.0-m2)*c + m2
	t.matrix[6] = mn*c1 + l*s
	t.matrix[8] = nl*c1 + m*s
	t.matrix[9] = mn*c1 - l*s
	t.matrix[10] = (1.0-n2)*c + n2
	return t
}

func lookAt(eye, gaze, up Vec3) Matrix {
	meye := eye.mult(-1)
	tv := translate(meye)

	// t axis = e-g
	t := eye.sub(gaze)
	// r axis = u cross t
	r := up.cross(t)
	// s axis = t cross r
	s := t.cross(r)

	rv := Matrix{}
	rv.loadIdentity()

	rlen := r.length()
	rv.matrix[0] = r[0] / rlen
	rv.matrix[4] = r[1] / rlen
	rv.matrix[8] = r[2] / rlen

	slen := s.length()
	rv.matrix[1] = s[0] / slen
	rv.matrix[5] = s[1] / slen
	rv.matrix[9] = s[2] / slen

	tlen := t.length()
	rv.matrix[2] = t[0] / tlen
	rv.matrix[6] = t[1] / tlen
	rv.matrix[10] = t[2] / tlen

	return rv.mult(tv)
}

func orthogonal(left, right, bottom, top, zNear, zFar float32) Matrix {
	dx := (right - left)
	dy := (top - bottom)
	dz := (zFar - zNear)

	t := Matrix{}
	if dx != 0.0 && dy != 0.0 && dz != 0.0 {
		t.loadIdentity()
		t.matrix[0] = 2.0 / dx
		t.matrix[5] = 2.0 / dy
		t.matrix[10] = -2.0 / dz
		t.matrix[12] = -(right + left) / dx
		t.matrix[13] = -(top + bottom) / dy
		t.matrix[14] = -(zFar + zNear) / dz
	}
	return t
}

func frustum(left, right, bottom, top, zNear, zFar float32) Matrix {
	dx := (right - left)
	dy := (top - bottom)
	dz := (zFar - zNear)

	t := Matrix{}
	if dx != 0.0 && dy != 0.0 && dz != 0.0 {
		t.loadIdentity()
		t.matrix[0] = 2.0 * zNear / dx
		t.matrix[5] = 2.0 * zNear / dy
		t.matrix[8] = (left + right) / 2
		t.matrix[9] = (top + bottom) / 2
		t.matrix[10] = -(zFar + zNear) / dz
		t.matrix[11] = -1
		t.matrix[14] = -2 * zFar * zNear / dz
		t.matrix[15] = 0
	}
	return t
}

func perspective(fovy, aspect, zNear, zFar float32) Matrix {
	dz := (zFar - zNear)

	t := Matrix{}
	if dz != 0.0 {
		t.loadIdentity()
		t.matrix[5] = 1 / tan(fovy*.5)
		t.matrix[0] = t.matrix[5] / aspect
		t.matrix[10] = -(zFar + zNear) / dz
		t.matrix[11] = -1
		t.matrix[14] = -2 * zFar * zNear / dz
		t.matrix[15] = 0
	}
	return t
}
