package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 600
const windowHeight = 480

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
	fmt.Println("OpenGL version:\t", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("GLSL version:\t", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("GLFW version:\t", glfw.GetVersionString())
	fmt.Println()

	glfw.SwapInterval(1)
	// Configure the vertex and fragment shaders
	var vertexShader = readFile("./chapter_4/point.vert")
	var fragmentShader = readFile("./chapter_4/point.frag")
	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
