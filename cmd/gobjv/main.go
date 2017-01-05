package main

import (
	"flag"
	"fmt"
	"github.com/zovt/go-obj-renderer/pkg/graphics"
	"github.com/zovt/go-obj-renderer/pkg/objs"
	"github.com/zovt/go-obj-renderer/pkg/web"
)

func main() {
	var path = flag.String("path", "", "The path of the obj file")
	var fp = flag.String("frag", "shaders/simple.glslf", "The fragment shader")
	var vp = flag.String("vert", "shaders/simple.glslv", "The vertex shader")
	flag.Parse()

	if *path == "" {
		fmt.Println("You must specify -path to an .obj file")
		return
	}

	// Start web server
	go web.Start()
	defer web.Close()

	// TODO: Implement full obj spec
	obj := objs.Parse(*path)

	graphics.Init()
	defer graphics.Close()
	graphics.LoadShaders(*vp, *fp)
	graphics.Render(obj)
}