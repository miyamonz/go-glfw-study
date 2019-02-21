package main

import (
	"fmt"
)

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

func (m1 *Matrix) mult(m2 *Matrix) Matrix {
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
