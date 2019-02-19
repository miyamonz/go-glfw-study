package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

func newProgram(vSrc, fSrc string) (uint32, error) {
	vShader, err := compileShader(vSrc, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fShader, err := compileShader(fSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)

	gl.BindAttribLocation(program, 0, gl.Str("position\x00"))
	gl.BindFragDataLocation(program, 0, gl.Str("fragment\x00"))

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: \n%v", log)
	}

	gl.DeleteShader(vShader)
	gl.DeleteShader(fShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile\n%v: %v", source, log)
	}

	return shader, nil
}
