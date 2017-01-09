package objs

import (
	"reflect"
	"testing"
)

var (
	tri, _  = Parse("../../objs/tri.obj")
	cube, _ = Parse("../../objs/cube.obj")
)

func TestVBOSize(t *testing.T) {
	if tri.VBOSize() != 84 {
		t.Errorf("Tri VBOSize incorrect.\nGot: %d\nExp: %d", tri.VBOSize(), 84)
	}

	if cube.VBOSize() != 224 {
		t.Errorf("Cube VBOSize incorrect.\nGot: %d\nExp: %d", cube.VBOSize(), 224)
	}
}

func TestVBOData(t *testing.T) {
	expTriData := []float32{
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		0.0, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
	}

	if !reflect.DeepEqual(tri.VBOData(), expTriData) {
		t.Error("Incorrect Tri data")
	}
}

func TestElementBufferData(t *testing.T) {
	triBuf := tri.ElementBufferData()
	cubeBuf := cube.ElementBufferData()

	if len(triBuf) != 3 {
		t.Error("Triangle EBO length incorrect")
	}

	if len(cubeBuf) != 36 {
		t.Error("Cube EBO length incorrect")
	}

	expTriBuf := []uint32{0, 1, 2}
	if !reflect.DeepEqual(expTriBuf, triBuf) {
		t.Error("Incorrect TriBuf data")
	}
}
