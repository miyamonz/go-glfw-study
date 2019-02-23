package main

import (
	"math"
)

func NewSphere() ShapeIndex {
	slices := 10
	stacks := 10

	vs := []Vertex{}
	for j := 0; j <= stacks; j++ {
		t := float32(j) / float32(stacks)
		y := cos(math.Pi * t)
		r := sin(math.Pi * t)

		for i := 0; i <= slices; i++ {
			s := float32(i) / float32(slices)
			z := r * cos(2*math.Pi*s)
			x := r * sin(2*math.Pi*s)

			pos := Vec3{x, y, z}
			color := Vec3{0.5, 0.8, 0.5}
			v := NewVertexN(pos, color, pos)
			vs = append(vs, v)

		}

	}
	index := []uint32{}
	for j := 0; j < stacks; j++ {
		k := uint32((slices + 1) * j)
		for i := 0; i < slices; i++ {
			k0 := k + uint32(i)
			k1 := k0 + 1
			k2 := k1 + uint32(slices)
			k3 := k2 + 1
			index = append(index, k0, k2, k3, k0, k3, k1)
		}
	}
	return NewSolidShapeIndex(vs, index)
}
