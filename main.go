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
	var nLoc = gl.GetUniformLocation(program, gl.Str("normalMatrix\x00"))

	var LposLoc = gl.GetUniformLocation(program, gl.Str("Lpos\x00"))
	var LambLoc = gl.GetUniformLocation(program, gl.Str("Lamb\x00"))
	var LdiffLoc = gl.GetUniformLocation(program, gl.Str("Ldiff\x00"))
	var LspecLoc = gl.GetUniformLocation(program, gl.Str("Lspec\x00"))

	w, h := window.GetSize()
	fw, fh := window.GetFramebufferSize()
	fmt.Println("aspect ratio: ", window.getAspect())
	fmt.Printf("width: %d, height: %d\n", w, h)
	fmt.Printf("frame buffer width: %d, frame buffer height: %d\n", fw, fh)

	shape := NewSphere()
	defer shape.Delete()

	//光源
	Lcount := 2
	Lpos := []Vec4{
		{0, 0, 5, 1},
		{8, 0, 1, 1},
	}
	Lamb := []Vec3{
		{0.2, 0.1, 0.1},
		{0.1, 0.1, 0.1},
	}
	Ldiff := []Vec3{
		{1.0, 0.5, 0.5},
		{0.9, 0.9, 0.9},
	}
	Lspec := []Vec3{
		{1.0, 0.5, 0.5},
		{0.9, 0.9, 0.9},
	}

	// draw
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	//カリング
	gl.FrontFace(gl.CCW)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.CULL_FACE)

	//デプスバッファ
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.DEPTH_TEST)

	glfw.SetTime(0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
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

		fovy := float32(1)
		aspect := window.getAspect()
		projection := perspective(fovy, aspect, 1, 100)

		normalMat := modelView.getNormalMat()
		gl.UniformMatrix4fv(mvLoc, 1, false, &modelView.data()[0])
		gl.UniformMatrix4fv(pLoc, 1, false, &projection.data()[0])
		gl.UniformMatrix4fv(nLoc, 1, false, &normalMat[0])
		shape.Draw()

		//二つ目
		modelView2 := modelView.mult(translate(Vec3{0, 0, 3}))
		normalMat2 := modelView2.getNormalMat()
		gl.UniformMatrix4fv(mvLoc, 1, false, &modelView2.data()[0])
		gl.UniformMatrix3fv(nLoc, 1, false, &normalMat2[0])
		shape.Draw()

		//光源

		for i := 0; i < Lcount; i++ {
			viewLpos := view.mul4x1(Lpos[i])
			gl.Uniform4fv(LposLoc+int32(i), 1, &viewLpos[0])
			gl.Uniform3fv(LambLoc, int32(Lcount), &Lamb[0][0])
			gl.Uniform3fv(LdiffLoc, int32(Lcount), &Ldiff[0][0])
			gl.Uniform3fv(LspecLoc, int32(Lcount), &Lspec[0][0])
		}
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
