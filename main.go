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
	window := initGLFW(windowWidth, windowHeight)
	defer glfw.Terminate()
	defer window.Destroy()

	//init gl
	//windowがmakecontextしないとむり
	if err := gl.Init(); err != nil {
		panic(err)
	}

	printDetail()

	//new program
	// Configure the vertex and fragment shaders
	var vertexShader = readFile("./shaders/point.vert")
	var fragmentShader = readFile("./shaders/point.frag")
	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	// buffer
	var rectPoints = []Vertex{
		{-0.5, -0.5},
		{0.5, -0.5},
		{0.5, 0.5},
		{-0.5, 0.5},
	}
	rect := NewShape(rectPoints)
	defer rect.Delete()

	// draw
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {


		draw(window, program, &rect)
	}
}

type Drawer interface {
	Draw()
}

func draw(window *glfw.Window, program uint32, drawer Drawer) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(program)

	drawer.Draw()

	window.SwapBuffers()
	glfw.PollEvents()
}

func initGLFW(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

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
	glfw.SwapInterval(1)
	return window
}

func printDetail() {
	fmt.Println("OpenGL version:\t", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("GLSL version:\t", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("GLFW version:\t", glfw.GetVersionString())
	fmt.Println()
}
