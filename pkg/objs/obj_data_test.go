package objs

import (
	"fmt"
	"testing"
)

func TestVertString(t *testing.T) {
	vert1 := Vertex{0, 0, 0, 1, nil, 0, 0}
	vert2 := Vertex{1, 2, 1, 1, nil, 0, 0}
	vertEmpty := Vertex{}

	v1ExpStr := "[x: 0.000000, y: 0.000000, z: 0.000000, w: 1.000000, faces: [], nID: 0, tID: 0]"
	v2ExpStr := "[x: 1.000000, y: 2.000000, z: 1.000000, w: 1.000000, faces: [], nID: 0, tID: 0]"
	vMtExpStr := "[x: 0.000000, y: 0.000000, z: 0.000000, w: 0.000000, faces: [], nID: 0, tID: 0]"

	if vert1.String() != v1ExpStr {
		t.Error("String conversion failed for vert1")
	}

	if vert2.String() != v2ExpStr {
		t.Error("String conversion failed for vert2")
	}

	if vertEmpty.String() != vMtExpStr {
		t.Error("String conversion failed for vertEmpty")
	}

	if fmt.Sprintf("%s", vert1) != v1ExpStr {
		t.Error("Sprintf failed for vert1")
	}

	if fmt.Sprintf("%s", vert2) != v2ExpStr {
		t.Error("Sprintf failed for vert2")
	}

	if fmt.Sprintf("%s", vertEmpty) != vMtExpStr {
		t.Error("Sprintf failed for vertEmpty")
	}
}

func TestFaceString(t *testing.T) {
	mtFace := Face{}
	face := Face{[]uint32{1, 2, 3}, []uint32{}, []uint32{}}

	mtFaceExpStr := "vIds: [], tIds: [], nIds: []"
	faceExpStr := "vIds: [1 2 3], tIds: [], nIds: []"

	if mtFace.String() != mtFaceExpStr {
		t.Error("String conversion for empty Face failed")
	}

	if face.String() != faceExpStr {
		t.Error("String converison failed for face")
	}
}

func TestObjDataString(t *testing.T) {
	odEmpty := ObjData{}
	odMtExpStr := "Vertices: [],\nFaces: [],\nTexture Coordinates: [],\nNormals: []\n"

	od1 := ObjData{[]Vertex{Vertex{1, 2, 1, 1, nil, 0, 0}}, []Face{}, []TexCoords{}, []Normal{}}
	od1ExpStr := "Vertices: [[x: 1.000000, y: 2.000000, z: 1.000000, w: 1.000000, faces: [], nID: 0, tID: 0]],\nFaces: [],\nTexture Coordinates: [],\nNormals: []\n"

	if odEmpty.String() != odMtExpStr {
		t.Error("Empty ObjData string not correct")
	}

	if od1.String() != od1ExpStr {
		t.Error("od1 string not correct")
	}
}
