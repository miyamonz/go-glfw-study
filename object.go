package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Object struct {
	vao, vbo, ibo uint32
}

func createVAO() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return vao
}
func createVBO(data []Vec3) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	size := len(data) * len(data[0])
	gl.BufferData(gl.ARRAY_BUFFER, size*4, gl.Ptr(data), gl.STATIC_DRAW)
	return vbo
}

func NewObject(points []Vec3) Object {
	vcount := len(points)
	if vcount == 0 {
		panic("points length is zero")
	}
	size := int32(len((points)[0]))
	obj := Object{
		vao: createVAO(),
		vbo: createVBO(points),
	}

	gl.VertexAttribPointer(0, size, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	return obj
}
func NewObjectWithIndex(points []Vec3, indexes []uint32) Object {

	obj := NewObject(points)

	gl.GenBuffers(1, &obj.ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, obj.ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indexes)*4, gl.Ptr(indexes), gl.STATIC_DRAW)

	return obj
}

func (obj *Object) Delete() {
	defer gl.DeleteBuffers(1, &obj.vao)
	defer gl.DeleteBuffers(1, &obj.vbo)
	defer gl.DeleteBuffers(1, &obj.ibo)
}

func (obj *Object) bind() {
	gl.BindVertexArray(obj.vao)
}

type Shape struct {
	object Object
	vcount int
}

func NewShape(points []Vec3) Shape {
	obj := NewObject(points)
	shape := Shape{
		object: obj,
		vcount: len(points),
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

func NewShapeIndex(points []Vec3, indexes []uint32) ShapeIndex {
	obj := NewObjectWithIndex(points, indexes)
	shape := Shape{
		object: obj,
		vcount: len(points),
	}
	shapei := ShapeIndex{
		Shape:      shape,
		indexcount: len(indexes),
		indexptr:   &indexes,
	}

	return shapei
}
func (s *ShapeIndex) Draw() {
	s.object.bind()
	gl.DrawElements(gl.LINES, int32(s.indexcount), gl.UNSIGNED_INT, gl.PtrOffset(0))
}
