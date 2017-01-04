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
	var nums [4]float64

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			nums[i] = v
		}
	}

	return Vertex{nums[0], nums[1], nums[2], nums[3]}
}

func parseNormal(args []string) Normal {
	var nums [3]float64

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			nums[i] = v
		}
	}

	return Normal{nums[0], nums[1], nums[2]}
}

func parseTexture(args []string) TexCoords {
	var nums [2]float64

	for i, s := range args {
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			nums[i] = v
		}
	}

	return TexCoords{nums[0], nums[1]}
}

func parseFace(args []string) Face {
	var vIds []uint
	var tIds []uint
	var nIds []uint

	for _, s := range args {
		sp := strings.Split(s, "/")
		v, err := strconv.ParseUint(sp[0], 10, 64)
		check(err)

		t, _ := strconv.ParseUint(sp[1], 10, 64)
		check(err)

		n, _ := strconv.ParseUint(sp[2], 10, 64)
		check(err)
		vIds = append(vIds, uint(v))
		tIds = append(tIds, uint(t))
		nIds = append(nIds, uint(n))
	}

	return Face{vIds, tIds, nIds}
}

func Parse(path string) ObjData {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	obj := ObjData{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) < 2 {
			continue
		}

		args := line[1:]
		switch line[0] {
		case "#":
		case "v":
			obj.Vertices = append(obj.Vertices, parseVertex(args))
		case "vn":
			obj.Normals = append(obj.Normals, parseNormal(args))
		case "vt":
			obj.TexCoords = append(obj.TexCoords, parseTexture(args))
		case "f":
			obj.Faces = append(obj.Faces, parseFace(args))
		}
	}

	check(scanner.Err())

	return obj
}
