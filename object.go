package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

type Vertex struct {
	position Vec3
	color    Vec3
}

func NewVertex(position, color Vec3) Vertex {
	return Vertex{
		position: position,
		color:    color,
	}
}

type Object struct {
	vao, vbo, ibo uint32
}

func createVAO() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return vao
}
func createVBO(data []Vertex) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	num := len(data)
	vertexSize := int(unsafe.Sizeof(data[0]))
	gl.BufferData(gl.ARRAY_BUFFER, num*vertexSize, gl.Ptr(data), gl.STATIC_DRAW)
	return vbo
}
func createIBO(index []uint32) uint32 {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	// 4はsizeof(uint32)ということ. 32bitは4byte
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(index)*4, gl.Ptr(index), gl.STATIC_DRAW)
	return ibo
}

func NewObject(points []Vertex) Object {
	vcount := len(points)
	if vcount == 0 {
		panic("points length is zero")
	}
	obj := Object{
		vao: createVAO(),
		vbo: createVBO(points),
	}

	vertexSize := int32(unsafe.Sizeof(points[0]))

	sizePos := int32(len(points[0].position))
	offsetPos := unsafe.Offsetof(points[0].position)
	//func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer)
	gl.VertexAttribPointer(0, sizePos, gl.FLOAT, false, vertexSize, gl.PtrOffset(int(offsetPos)))
	gl.EnableVertexAttribArray(0)

	sizeColor := int32(len(points[0].color))
	offsetColor := unsafe.Offsetof(points[0].color)
	gl.VertexAttribPointer(1, sizeColor, gl.FLOAT, false, vertexSize, gl.PtrOffset(int(offsetColor)))
	gl.EnableVertexAttribArray(1)

	return obj
}
func NewObjectWithIndex(points []Vertex, indexes []uint32) Object {
	obj := NewObject(points)
	obj.ibo = createIBO(indexes)
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

func NewShape(points []Vertex) Shape {
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

func NewShapeIndex(points []Vertex, indexes []uint32) ShapeIndex {
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
