package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 600
const windowHeight = 480
var points = []float32{
	-0.5, -0.5,
	0.5, -0.5,
	0.5, 0.5,
	-0.5, 0.5,
}
func init() {
	runtime.LockOSThread()
}
func main() {
	//init glfw
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// make an application window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Hello", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	//init gl
	//windowがmakecontextしないとむり
	if err := gl.Init(); err != nil {
		panic(err)
	}

	//log status
	fmt.Println("OpenGL version:\t", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("GLSL version:\t", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("GLFW version:\t", glfw.GetVersionString())
	fmt.Println()

	//new program
	glfw.SwapInterval(1)
	// Configure the vertex and fragment shaders
	var vertexShader = readFile("./chapter_4/point.vert")
	var fragmentShader = readFile("./chapter_4/point.frag")
	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	// buffer
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(points) * 4, gl.Ptr(points), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	defer gl.DeleteBuffers(1, &vao)
	defer gl.DeleteBuffers(1, &vbo)



	// draw
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.LINE_LOOP, 0,4 )

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
