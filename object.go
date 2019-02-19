package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Vertex [2]float32

type Object struct {
	vao, vbo uint32
}

func NewObject(points []Vertex) Object {
	vcount := len(points)
	if vcount == 0 {
		panic("points length is zero")
	}
	size := int32(len((points)[0]))
	obj := Object{}
	gl.GenVertexArrays(1, &obj.vao)
	gl.BindVertexArray(obj.vao)

	gl.GenBuffers(1, &obj.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, obj.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, vcount*int(size)*4, gl.Ptr(points), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, size, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	return obj
}

func (obj *Object) Delete() {
	defer gl.DeleteBuffers(1, &obj.vao)
	defer gl.DeleteBuffers(1, &obj.vbo)
}

func (obj *Object) bind() {
	gl.BindVertexArray(obj.vao)
}

type Shape struct {
	object Object
	vcount int
}

func NewShape(points []Vertex) Shape {
	obj := NewObject(points)
	shape := Shape{
		obj,
		len(points),
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
