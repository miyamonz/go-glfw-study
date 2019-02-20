package main

import (
	"math"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	*glfw.Window
	size [2]float32
	title string
	scale float32
}

func NewWindow(w,h int) (*Window, error) {
	win,err :=  initGLFW(w,h)
	if err != nil {
		return nil,err
	}
	ret := &Window{
		Window: win,
		size: [2]float32{float32(w), float32(h)},
		title: "",
		scale: 100,
	}
	win.SetSizeCallback( func(win *glfw.Window, w,h int ) {
		ret.resize(w,h)

	} )
	return ret, nil
}

func (win *Window) Destroy() {
	win.Window.Destroy()
	//もしもmulti winodwならこれはおかしいがまあいいや
	glfw.Terminate()
}

func (win *Window) resize(width,height int) {
	// Retina Display must use FrameBufferSize
	fw, fh := win.GetFramebufferSize()
	gl.Viewport(0, 0, int32(fw), int32(fh))
	// fmt.Printf("resize called. width: %d, height: %d, aspect: %f\n", int32(width), int32(height), aspect)
	win.size[0] = float32(width)
	win.size[1] = float32(height)
}
func (win *Window) getAspect() float32 {
	return win.size[0] / win.size[1]
}
func (win *Window) update() {
	a := float32(50 * math.Cos(sec() * 2 * math.Pi))
	win.scale = 100 + a
}


func initGLFW(width, height int) (*glfw.Window,  error) {
	if err := glfw.Init(); err != nil {
		return nil,err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// make an application window
	window, err := glfw.CreateWindow(width, height, "Hello", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	return window, nil
}
