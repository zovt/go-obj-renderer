package main

import (
	"flag"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/zovt/go-obj-renderer/pkg/graphics"
	"github.com/zovt/go-obj-renderer/pkg/objs"
	"github.com/zovt/go-obj-renderer/pkg/web"
	"runtime"
)

func main() {
	// Lock the main OS thread to prevent some weird GLFW errors
	runtime.LockOSThread()

	var path = flag.String("path", "", "The path of the obj file")
	var fp = flag.String("frag", "shaders/simple.glslf", "The fragment shader")
	var vp = flag.String("vert", "shaders/simple.glslv", "The vertex shader")
	flag.Parse()

	if *path == "" {
		fmt.Println("You must specify -path to an .obj file")
		return
	}

	ch := make(chan func(a, b mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3))

	go func() {
		err := web.Start(ch)
		if err != nil {
			fmt.Println("Could not start web server")
			panic(err)
		}
	}()

	obj, err := objs.Parse(*path)
	if err != nil {
		panic(err)
	}

	graphics.Init()
	defer graphics.Close()
	graphics.LoadShaders(*vp, *fp)
	graphics.Render(obj, ch)
}
