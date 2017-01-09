package objs

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tri := Parse("../../objs/tri.obj")
	cube := Parse("../../objs/cube.obj")

	cubeStr := "Vertices: [[x: 0.000000, y: 0.000000, z: 0.000000, w: 0.000000, faces: [1 2 3 4 9 10], nID: 1, tID: 0] [x: 0.000000, y: 0.000000, z: 1.000000, w: 0.000000, faces: [4 10 11 12], nID: 2, tID: 0] [x: 0.000000, y: 1.000000, z: 0.000000, w: 0.000000, faces: [2 3 5 6], nID: 3, tID: 0] [x: 0.000000, y: 1.000000, z: 1.000000, w: 0.000000, faces: [3 4 6 12], nID: 4, tID: 0] [x: 1.000000, y: 0.000000, z: 0.000000, w: 0.000000, faces: [1 7 8 9], nID: 5, tID: 0] [x: 1.000000, y: 0.000000, z: 1.000000, w: 0.000000, faces: [8 9 10 11], nID: 6, tID: 0] [x: 1.000000, y: 1.000000, z: 0.000000, w: 0.000000, faces: [1 2 5 7], nID: 0, tID: 0] [x: 1.000000, y: 1.000000, z: 1.000000, w: 0.000000, faces: [5 6 7 8 11 12], nID: 0, tID: 0]],\nFaces: [vIds: [1 7 5], tIds: [0 0 0], nIds: [2 2 2] vIds: [1 3 7], tIds: [0 0 0], nIds: [2 2 2] vIds: [1 4 3], tIds: [0 0 0], nIds: [6 6 6] vIds: [1 2 4], tIds: [0 0 0], nIds: [6 6 6] vIds: [3 8 7], tIds: [0 0 0], nIds: [3 3 3] vIds: [3 4 8], tIds: [0 0 0], nIds: [3 3 3] vIds: [5 7 8], tIds: [0 0 0], nIds: [5 5 5] vIds: [5 8 6], tIds: [0 0 0], nIds: [5 5 5] vIds: [1 5 6], tIds: [0 0 0], nIds: [4 4 4] vIds: [1 6 2], tIds: [0 0 0], nIds: [4 4 4] vIds: [2 6 8], tIds: [0 0 0], nIds: [1 1 1] vIds: [2 8 4], tIds: [0 0 0], nIds: [1 1 1]],\nTexture Coordinates: [],\nNormals: [{0 0 1 0} {0 0 -1 0} {0 1 0 0} {0 -1 0 0} {1 0 0 0} {-1 0 0 0}]\n"

	triVertices := []Vertex{Vertex{-0.5, -0.5, 0, 0, []uint32{1}, 0, 0}, Vertex{0.5, -0.5, 0, 0, []uint32{1}, 0, 0}, Vertex{0, 0.5, 0, 0, []uint32{1}, 0, 0}}
	triFaces := []Face{Face{[]uint32{1, 2, 3}, []uint32{0, 0, 0}, []uint32{0, 0, 0}}}

	if !reflect.DeepEqual(triVertices, tri.Vertices) {
		t.Errorf("Incorrect vertices.\nGot: %s\nExp: %s", tri.Vertices, triVertices)
	}

	if !reflect.DeepEqual(triFaces, tri.Faces) {
		t.Errorf("Incorrect faces.\nGot: %v\nExp: %v", tri.Faces, triFaces)
	}

	if cube.String() != cubeStr {
		t.Errorf("Cube string not correct")
	}
}
