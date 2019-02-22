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
	window, err := NewWindow(600, 480)
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

	var mvLoc = gl.GetUniformLocation(program, gl.Str("modelview\x00"))
	var pLoc = gl.GetUniformLocation(program, gl.Str("projection\x00"))

	w, h := window.GetSize()
	fw, fh := window.GetFramebufferSize()
	fmt.Println("aspect ratio: ", window.getAspect())
	fmt.Printf("width: %d, height: %d\n", w, h)
	fmt.Printf("frame buffer width: %d, frame buffer height: %d\n", fw, fh)

	cube := NewSolidCube()
	defer cube.Delete()

	// draw
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	//カリング
	gl.FrontFace(gl.CCW)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.CULL_FACE)

	glfw.SetTime(0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)
		window.update()

		//scale matrix
		s := window.scale * 2
		scaleMat := scale(Vec3{s, s, s})

		//translate matrix
		location := window.location
		trans := Vec3{location[0], location[1], 0}
		transMat := translate(trans)
		view := lookAt(
			Vec3{3, 4, 5}, //eye
			Vec3{0, 0, 0}, //gaze
			Vec3{0, 1, 0},
		)
		rotMat := rotate(float32(glfw.GetTime()), Vec3{0, 1, 0})

		model := transMat.mult(rotMat).mult(scaleMat)
		modelView := view.mult(model)
		modelView2 := modelView.mult(translate(Vec3{0, 0, 3}))

		fovy := float32(1)
		aspect := window.getAspect()
		projection := perspective(fovy, aspect, 1, 100)
		gl.UniformMatrix4fv(mvLoc, 1, false, &modelView.data()[0])
		gl.UniformMatrix4fv(pLoc, 1, false, &projection.data()[0])
		cube.Draw()

		//二つ目
		gl.UniformMatrix4fv(mvLoc, 1, false, &modelView2.data()[0])
		cube.Draw()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

type IWindow interface {
	SwapBuffers()
}
type Drawer interface {
	Draw()
}

func printDetail() {
	fmt.Println("OpenGL version:\t", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("GLSL version:\t", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("GLFW version:\t", glfw.GetVersionString())
}
