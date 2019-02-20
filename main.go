package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	//init glfw
	window,err := NewWindow(600, 480)
	if err != nil {
		panic(err)
	}
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

	var sizeLoc = gl.GetUniformLocation(program, gl.Str("size\x00"))
	var scaleLoc = gl.GetUniformLocation(program, gl.Str("scale\x00"))
	var locationLoc = gl.GetUniformLocation(program, gl.Str("location\x00"))

	w, h := window.GetSize()
	fw, fh := window.GetFramebufferSize()
	fmt.Println("aspect ratio: ", window.getAspect())
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
		window.update()
		gl.Uniform1f(scaleLoc, window.scale)
		gl.Uniform2f(locationLoc, window.location[0], window.location[1])

		fw, fh := window.GetSize()
		gl.Uniform2f(sizeLoc, float32(fw), float32(fh))

		draw(window, program, &rect)
	}
}

type IWindow interface {
	SwapBuffers()
}
type Drawer interface {
	Draw()
}

func draw(window IWindow, program uint32, drawer Drawer) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(program)

	drawer.Draw()

	window.SwapBuffers()
	glfw.PollEvents()
}


func printDetail() {
	fmt.Println("OpenGL version:\t", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("GLSL version:\t", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("GLFW version:\t", glfw.GetVersionString())
}

