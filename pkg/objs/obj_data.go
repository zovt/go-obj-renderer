package objs

import "fmt"

type Vertex struct {
	x     float32
	y     float32
	z     float32
	w     float32
	faces []uint32
	nID   uint32
	tID   uint32
}

func (v Vertex) String() string {
	return fmt.Sprintf("[x: %f, y: %f, z: %f, w: %f, faces: %v, nID: %d, tID: %d]", v.x, v.y, v.z, v.w, v.faces, v.nID, v.tID)
}

type Face struct {
	vIds []uint32
	tIds []uint32
	nIds []uint32
}

func (f Face) String() string {
	return fmt.Sprintf("vIds: %v, tIds: %v, nIds: %v", f.vIds, f.tIds, f.nIds)
}

type TexCoords struct {
	u float32
	v float32
}

type Normal struct {
	x   float32
	y   float32
	z   float32
	vID uint32
}

type ObjData struct {
	Vertices  []Vertex
	Faces     []Face
	TexCoords []TexCoords
	Normals   []Normal
}

func (od ObjData) String() string {
	return fmt.Sprintf("Vertices: %v,\nFaces: %v,\nTexture Coordinates: %v,\nNormals: %v\n",
		od.Vertices,
		od.Faces,
		od.TexCoords,
		od.Normals)
}
