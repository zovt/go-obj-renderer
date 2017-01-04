package objs

import "fmt"

type Vertex struct {
	x float64
	y float64
	z float64
	w float64
}

func (v Vertex) String() string {
	return fmt.Sprintf("[x: %f, y: %f, z: %f, w: %f]", v.x, v.y, v.z, v.w)
}

type Face struct {
	vIds []uint
	tIds []uint
	nIds []uint
}

func (f Face) String() string {
	return fmt.Sprintf("vIds: %v, tIds: %v, nIds: %v", f.vIds, f.tIds, f.nIds)
}

type TexCoords struct {
	u float64
	v float64
}

type Normal struct {
	x float64
	y float64
	z float64
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
