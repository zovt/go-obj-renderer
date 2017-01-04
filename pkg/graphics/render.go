package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/zovt/go-obj-renderer/pkg/objs"
	"io/ioutil"
	"runtime"
	"strings"
)

const w = 800
const h = 600

var window *glfw.Window

func Init() {
	// Lock OS Thread for GLFW events
	runtime.LockOSThread()

	if e := glfw.Init(); e != nil {
		panic(e)
	}

	// Hint GLFW window
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var e error
	window, e = glfw.CreateWindow(w, h, "OBJ Viewer", nil, nil)
	if e != nil {
		panic(e)
	}

	window.MakeContextCurrent()

	// Init GLOW
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func LoadShaders(vp, fp string) {
	vs, e := loadShader(vp, gl.VERTEX_SHADER)
	if e != nil {
		panic(e)
	}

	fs, e := loadShader(fp, gl.FRAGMENT_SHADER)
	if e != nil {
		panic(e)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vs)
	gl.AttachShader(prog, fs)
	gl.LinkProgram(prog)

	var status int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(prog, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(prog, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	gl.UseProgram(prog)
}

func loadShader(p string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	source, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	sourceStr := string(source) + "\x00"

	csources, free := gl.Strs(sourceStr)
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

		return 0, fmt.Errorf("failed to compile \n%s\n%v", string(source), log)
	}

	return shader, nil
}

func Render(obj objs.ObjData) {
	// configure VAO and VBO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(obj.Vertices)*4, gl.Ptr(obj.Vertices), gl.STATIC_DRAW)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func Close() {
	glfw.Terminate()
}
