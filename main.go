package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 600
const windowHeight = 480

var aspect = float32(windowWidth) / float32(windowHeight)
var scale float32 = 100.0
var size = [2]float32{float32(windowWidth), float32(windowHeight)}

func init() {
	runtime.LockOSThread()
}
func main() {
	//init glfw
	window := initGLFW(windowWidth, windowHeight)
	window.SetSizeCallback(resize)
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

	var aspectLoc = gl.GetUniformLocation(program, gl.Str("aspect\x00"))
	var sizeLoc = gl.GetUniformLocation(program, gl.Str("size\x00"))
	var scaleLoc = gl.GetUniformLocation(program, gl.Str("scale\x00"))

	w, h := window.GetSize()
	fw, fh := window.GetFramebufferSize()
	fmt.Println("aspect ratio: ", aspect)
	fmt.Printf("width: %d, height: %d\n", w, h)
	fmt.Printf("frame buffer width: %d, frame buffer height: %d\n", fw, fh)

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
	gl.Viewport(0, 0, int32(fw), int32(fh))

	for !window.ShouldClose() {
		gl.Uniform1f(aspectLoc, aspect)
		gl.Uniform1f(scaleLoc, scale)

		fw, fh := window.GetSize()
		gl.Uniform2f(sizeLoc, float32(fw), float32(fh))

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

func resize(w *glfw.Window, width, height int) {
	// Retina Display must use FrameBufferSize
	fw, fh := w.GetFramebufferSize()
	gl.Viewport(0, 0, int32(fw), int32(fh))
	// fmt.Printf("resize called. width: %d, height: %d, aspect: %f\n", int32(width), int32(height), aspect)
	aspect = float32(width) / float32(height)
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
}
