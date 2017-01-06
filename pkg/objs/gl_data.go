package objs

type GLData interface {
	VBOData() []float32
	VBOSize() int
	ElementBufferData() []uint
}

func (obj ObjData) VBOSize() int {
	return (len(obj.Vertices) * 7) * 4
}

func (obj ObjData) VBOData() []float32 {
	data := make([]float32, len(obj.Vertices)*7)

	offset := 0
	for _, v := range obj.Vertices {
		data[offset+0] = v.x
		data[offset+1] = v.y
		data[offset+2] = v.z
		data[offset+3] = 1.0
		offset += 4
		data[offset+0] = 0
		data[offset+1] = 0
		data[offset+2] = 0
		offset += 3
	}

	return data
}

func (obj ObjData) ElementBufferData() []uint32 {
	indices := make([]uint32, len(obj.Faces)*3)

	offset := 0
	for _, f := range obj.Faces {
		for _, n := range f.vIds {
			indices[offset] = n
			offset += 1
		}
	}

	return indices
}
