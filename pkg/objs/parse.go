package objs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseVertex(args []string) Vertex {
	var nums [4]float32

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			nums[i] = float32(v)
		}
	}

	return Vertex{nums[0], nums[1], nums[2], nums[3], []uint32{}, 0, 0}
}

func parseNormal(args []string) Normal {
	var nums [3]float32

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			nums[i] = float32(v)
		}
	}

	return Normal{nums[0], nums[1], nums[2], 0}
}

func parseTexture(args []string) TexCoords {
	var nums [2]float32

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			nums[i] = float32(v)
		}
	}

	return TexCoords{nums[0], nums[1]}
}

func parseFace(obj ObjData, idx uint32, args []string) Face {
	var vIds []uint32
	var tIds []uint32
	var nIds []uint32

	for _, s := range args {
		sp := strings.Split(s, "/")
		v, err := strconv.ParseUint(sp[0], 10, 32)
		if err != nil {
			v = 0
		}

		vIds = append(vIds, uint32(v))
		obj.Vertices[v-1].faces = append(obj.Vertices[v-1].faces, idx+1)

		if len(sp) == 1 {
			tIds = append(tIds, 0)
			nIds = append(nIds, 0)
			continue
		}

		t, err := strconv.ParseUint(sp[1], 10, 32)
		if err != nil {
			t = 0
		}

		tIds = append(tIds, uint32(t))
		if t != 0 {
			obj.Vertices[t-1].tID = uint32(t)
		}

		if len(sp) == 2 {
			nIds = append(nIds, 0)
			continue
		}

		n, err := strconv.ParseUint(sp[2], 10, 32)
		if err != nil {
			t = 0
		}

		nIds = append(nIds, uint32(n))
		if n != 0 {
			obj.Vertices[n-1].nID = uint32(n)
		}
	}

	return Face{vIds, tIds, nIds}
}

func Parse(path string) ObjData {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	obj := ObjData{}

	scanner := bufio.NewScanner(file)
	var fIdx uint32
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) < 2 {
			continue
		}

		args := line[1:]
		switch line[0] {
		case "v":
			obj.Vertices = append(obj.Vertices, parseVertex(args))
		case "vn":
			obj.Normals = append(obj.Normals, parseNormal(args))
		case "vt":
			obj.TexCoords = append(obj.TexCoords, parseTexture(args))
		case "f":
			obj.Faces = append(obj.Faces, parseFace(obj, fIdx, args))
			fIdx++
		}
	}

	check(scanner.Err())

	return obj
}
