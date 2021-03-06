package renderer

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/gopherd/doge/math/tensor"

	"github.com/gopherd/three/driver/renderer/shader"
)

type openglRenderer struct {
}

func OpenGLRenderer() Renderer {
	return &openglRenderer{}
}

func (openglRenderer) Init(width, height int) error {
	if err := gl.Init(); err != nil {
		return err
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	return nil
}

func (openglRenderer) Viewport(x, y, w, h int32) {
	gl.Viewport(x, y, w, h)
}

func (openglRenderer) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func createShader(shaderType uint32, source string) (uint32, error) {
	var success int32
	var shaderId = gl.CreateShader(shaderType)
	var ptr = (*byte)(unsafe.Pointer(&source))
	gl.ShaderSource(shaderId, 1, &ptr, nil)
	gl.CompileShader(shaderId)
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &success)
	if success == 0 {
		return shaderId, nil
	}
	const size = 512
	var buf = make([]byte, size)
	var n int32
	gl.GetShaderInfoLog(shaderId, size, &n, &buf[0])
	return 0, errors.New(string(buf[:n]))
}

func (openglRenderer) CreateProgram(vshader, fshader string) (program Program, err error) {
	var (
		vshaderId uint32
		fshaderId uint32
	)
	vshaderId, err = createShader(gl.VERTEX_SHADER, vshader)
	if err != nil {
		return
	}
	fshaderId, err = createShader(gl.FRAGMENT_SHADER, fshader)
	if err != nil {
		return
	}
	var id = gl.CreateProgram()
	gl.AttachShader(id, vshaderId)
	gl.AttachShader(id, fshaderId)
	gl.DeleteShader(vshaderId)
	gl.DeleteShader(fshaderId)
	program = Program{
		Id:               id,
		VertextShaderId:  vshaderId,
		FragmentShaderId: fshaderId,
	}
	return
}

func (openglRenderer) ClearProgram(program Program) {
	gl.DeleteProgram(program.Id)
}

func (openglRenderer) LinkProgram(program uint32) error {
	gl.LinkProgram(program)
	var success int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success == 0 {
		return nil
	}
	const size = 512
	var buf = make([]byte, size)
	var n int32
	gl.GetProgramInfoLog(program, size, &n, &buf[0])
	return errors.New(string(buf[:n]))
}

func (openglRenderer) SetUniform(program uint32, name string, uniform shader.Uniform) {
	var location = gl.GetUniformLocation(program, (*byte)(unsafe.Pointer(&name)))
	switch value := uniform.(type) {
	case int:
		gl.Uniform1i(location, int32(value))
	case [2]int:
		gl.Uniform2i(location, int32(value[0]), int32(value[1]))
	case [3]int:
		gl.Uniform3i(location, int32(value[0]), int32(value[1]), int32(value[2]))
	case [4]int:
		gl.Uniform4i(location, int32(value[0]), int32(value[1]), int32(value[2]), int32(value[3]))
	case int8:
		gl.Uniform1i(location, int32(value))
	case [2]int8:
		gl.Uniform2i(location, int32(value[0]), int32(value[1]))
	case [3]int8:
		gl.Uniform3i(location, int32(value[0]), int32(value[1]), int32(value[2]))
	case [4]int8:
		gl.Uniform4i(location, int32(value[0]), int32(value[1]), int32(value[2]), int32(value[3]))
	case int16:
		gl.Uniform1i(location, int32(value))
	case [2]int16:
		gl.Uniform2i(location, int32(value[0]), int32(value[1]))
	case [3]int16:
		gl.Uniform3i(location, int32(value[0]), int32(value[1]), int32(value[2]))
	case [4]int16:
		gl.Uniform4i(location, int32(value[0]), int32(value[1]), int32(value[2]), int32(value[3]))
	case int32:
		gl.Uniform1i(location, int32(value))
	case [2]int32:
		gl.Uniform2i(location, int32(value[0]), int32(value[1]))
	case [3]int32:
		gl.Uniform3i(location, int32(value[0]), int32(value[1]), int32(value[2]))
	case [4]int32:
		gl.Uniform4i(location, int32(value[0]), int32(value[1]), int32(value[2]), int32(value[3]))
	case uint:
		gl.Uniform1ui(location, uint32(value))
	case [2]uint:
		gl.Uniform2ui(location, uint32(value[0]), uint32(value[1]))
	case [3]uint:
		gl.Uniform3ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]))
	case [4]uint:
		gl.Uniform4ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]), uint32(value[3]))
	case uint8:
		gl.Uniform1ui(location, uint32(value))
	case [2]uint8:
		gl.Uniform2ui(location, uint32(value[0]), uint32(value[1]))
	case [3]uint8:
		gl.Uniform3ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]))
	case [4]uint8:
		gl.Uniform4ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]), uint32(value[3]))
	case uint16:
		gl.Uniform1ui(location, uint32(value))
	case [2]uint16:
		gl.Uniform2ui(location, uint32(value[0]), uint32(value[1]))
	case [3]uint16:
		gl.Uniform3ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]))
	case [4]uint16:
		gl.Uniform4ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]), uint32(value[3]))
	case uint32:
		gl.Uniform1ui(location, uint32(value))
	case [2]uint32:
		gl.Uniform2ui(location, uint32(value[0]), uint32(value[1]))
	case [3]uint32:
		gl.Uniform3ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]))
	case [4]uint32:
		gl.Uniform4ui(location, uint32(value[0]), uint32(value[1]), uint32(value[2]), uint32(value[3]))
	case float32:
		gl.Uniform1f(location, value)
	case [2]float32:
		gl.Uniform2f(location, value[0], value[1])
	case tensor.Vector2[float32]:
		gl.Uniform2f(location, value.X(), value.Y())
	case [3]float32:
		gl.Uniform3f(location, value[0], value[1], value[2])
	case tensor.Vector3[float32]:
		gl.Uniform3f(location, value.X(), value.Y(), value.Z())
	case [4]float32:
		gl.Uniform4f(location, value[0], value[1], value[2], value[3])
	case tensor.Vector4[float32]:
		gl.Uniform4f(location, value.X(), value.Y(), value.Z(), value.W())
	case float64:
		gl.Uniform1d(location, value)
	case [2]float64:
		gl.Uniform2d(location, value[0], value[1])
	case tensor.Vector2[float64]:
		gl.Uniform2d(location, value.X(), value.Y())
	case [3]float64:
		gl.Uniform3d(location, value[0], value[1], value[2])
	case tensor.Vector3[float64]:
		gl.Uniform3d(location, value.X(), value.Y(), value.Z())
	case [4]float64:
		gl.Uniform4d(location, value[0], value[1], value[2], value[3])
	case tensor.Vector4[float64]:
		gl.Uniform4d(location, value.X(), value.Y(), value.Z(), value.W())
	case tensor.Matrix3[float32]:
		gl.UniformMatrix3fv(location, 1, false, &value[0])
	case tensor.Matrix3[float64]:
		gl.UniformMatrix3dv(location, 1, false, &value[0])
	case tensor.Matrix4[float32]:
		gl.UniformMatrix4fv(location, 1, false, &value[0])
	case tensor.Matrix4[float64]:
		gl.UniformMatrix4dv(location, 1, false, &value[0])
	default:
		panic(fmt.Sprintf("unsupported uniform type: %T", uniform))
	}
}
