package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"reflect"
	"unsafe"
)

type Vertex struct {
	position Vec3
	color    Vec3
	normal   Vec3
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
func registerAttr(loc uint32, v interface{}, name string) {
	VType := reflect.TypeOf(v)
	vertexSize := VType.Size()

	attr, ok := VType.FieldByName(name)
	if !ok {
		panic("failed VType FieldByName:")
	}
	size := attr.Type.Len()
	offset := attr.Offset
	gl.VertexAttribPointer(loc, int32(size), gl.FLOAT, false, int32(vertexSize), gl.PtrOffset(int(offset)))
	gl.EnableVertexAttribArray(loc)
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

	registerAttr(0, Vertex{}, "position")
	registerAttr(1, Vertex{}, "color")
	registerAttr(2, Vertex{}, "normal")

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
