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

func NewCube() ShapeIndex {
	var wireCubeIndex []uint32 = []uint32{
		1, 0,
		2, 7,
		3, 0,
		4, 7,
		5, 0,
		6, 7,
		1, 2,
		2, 3,
		3, 4,
		4, 5,
		5, 6,
		6, 1,
	}
	return NewShapeIndex(cubeVertex, wireCubeIndex)
}

func NewSolidCube() ShapeIndex {
	var corners []Vec3 = []Vec3{
		Vec3{-1.0, -1.0, -1.0},
		Vec3{-1.0, -1.0, 1.0},
		Vec3{-1.0, 1.0, 1.0},
		Vec3{-1.0, 1.0, -1.0},
		Vec3{1.0, 1.0, -1.0},
		Vec3{1.0, -1.0, -1.0},
		Vec3{1.0, -1.0, 1.0},
		Vec3{1.0, 1.0, 1.0},
	}
	red := Vec3{0.8, 0.1, 0.1}
	green := Vec3{0.1, 0.8, 0.1}
	blue := Vec3{0.1, 0.1, 0.8}
	yellow := Vec3{0.8, 0.8, 0.1}
	magenta := Vec3{0.8, 0.1, 0.8}
	cyan := Vec3{0.1, 0.8, 0.8}
	var cubeVertex []Vertex = []Vertex{
		// 左
		NewVertex(corners[0], green),
		NewVertex(corners[1], green),
		NewVertex(corners[2], green),
		NewVertex(corners[3], green),
		// 裏
		NewVertex(corners[5], magenta),
		NewVertex(corners[0], magenta),
		NewVertex(corners[3], magenta),
		NewVertex(corners[4], magenta),
		// 下
		NewVertex(corners[0], cyan),
		NewVertex(corners[5], cyan),
		NewVertex(corners[6], cyan),
		NewVertex(corners[1], cyan),
		// 右
		NewVertex(corners[6], blue),
		NewVertex(corners[5], blue),
		NewVertex(corners[4], blue),
		NewVertex(corners[7], blue),
		// 上
		NewVertex(corners[3], red),
		NewVertex(corners[2], red),
		NewVertex(corners[7], red),
		NewVertex(corners[4], red),
		// 前
		NewVertex(corners[1], yellow),
		NewVertex(corners[6], yellow),
		NewVertex(corners[7], yellow),
		NewVertex(corners[2], yellow),
	}
	var solidCubeIndex []uint32 = []uint32{
		0, 1, 2, 0, 2, 3, // 左
		4, 5, 6, 4, 6, 7, // 裏
		8, 9, 10, 8, 10, 11, //下
		12, 13, 14, 12, 14, 15, // 右
		16, 17, 18, 16, 18, 19, // 上
		20, 21, 22, 20, 22, 23, //前
	}
	return NewSolidShapeIndex(cubeVertex, solidCubeIndex)
}
