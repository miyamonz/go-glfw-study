package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shape struct {
	object Object
	vcount int
	mode   uint32
}

func NewShape(points []Vertex) Shape {
	obj := NewObject(points)
	shape := Shape{
		object: obj,
		vcount: len(points),
		mode:   gl.LINES,
	}

	return shape
}

func (s *Shape) Draw() {
	s.object.bind()
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(s.vcount))
}

func (shape *Shape) Delete() {
	shape.object.Delete()
}

type ShapeIndex struct {
	Shape
	indexcount int
	indexptr   *[]uint32
}

func NewShapeIndex(points []Vertex, indexes []uint32) ShapeIndex {
	obj := NewObjectWithIndex(points, indexes)
	shape := Shape{
		object: obj,
		vcount: len(points),
		mode:   gl.LINES,
	}
	shapei := ShapeIndex{
		Shape:      shape,
		indexcount: len(indexes),
		indexptr:   &indexes,
	}

	return shapei
}
func NewSolidShapeIndex(points []Vertex, indexes []uint32) ShapeIndex {
	shape := NewShapeIndex(points, indexes)
	shape.mode = gl.TRIANGLES
	return shape
}

func (s *ShapeIndex) Draw() {
	s.object.bind()
	gl.DrawElements(s.mode, int32(s.indexcount), gl.UNSIGNED_INT, gl.PtrOffset(0))
}

var cubeVertex []Vertex = []Vertex{
	NewVertex(Vec3{-1, -1, -1}, Vec3{0, 0, 0}),
	NewVertex(Vec3{-1, -1, 1}, Vec3{0, 0, 8}),
	NewVertex(Vec3{-1, 1, 1}, Vec3{0, .8, 0}),
	NewVertex(Vec3{-1, 1, -1}, Vec3{0, .8, .8}),
	NewVertex(Vec3{1, 1, -1}, Vec3{.8, 0, 0}),
	NewVertex(Vec3{1, -1, -1}, Vec3{.8, 0, .8}),
	NewVertex(Vec3{1, -1, 1}, Vec3{.8, .8, 0}),
	NewVertex(Vec3{1, 1, 1}, Vec3{.8, .8, .8}),
}
