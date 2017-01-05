package objs

type GLData interface {
	VBOData() []float32
	VBOSize() int
	ElementBufferData() []uint
}

func (obj ObjData) VBOSize() int {
	return (len(obj.Vertices) * 4) * 4
}

func (obj ObjData) VBOData() []float32 {
	vertices := make([]float32, len(obj.Vertices)*4)

	offset := 0
	for _, v := range obj.Vertices {
		vertices[offset+0] = v.x
		vertices[offset+1] = v.y
		vertices[offset+2] = v.z
		vertices[offset+3] = 1.0
		offset += 4
	}

	return vertices
}

func (obj ObjData) ElementBufferData() []uint32 {
	indices := make([]uint32, len(obj.Faces)*3)

	offset := 0
	for _, f := range obj.Faces {
		for _, n := range f.vIds {
			indices[offset] = n - 1
			offset += 1
		}
	}

	return indices
}
