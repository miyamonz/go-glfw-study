package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	*glfw.Window
	width, height int
	title string
}

func NewWindow(w,h int) (*Window, error) {
	win,err :=  initGLFW(w,h)
	if err != nil {
		return nil,err
	}
	ret := &Window{
		win,
		w,
		h,
		"",
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
	aspect = float32(width) / float32(height)
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
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Hello", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	return window, nil
}
