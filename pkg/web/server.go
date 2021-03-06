package web

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"net/http"
)

func zoom(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	displacement := from.Sub(at).Normalize()
	return from.Sub(displacement.Mul(dist)), at
}

func moveX(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	d := from.Sub(at)
	theta := float32(math.Atan(float64(d.X() / d.Z())))
	disp := (mgl32.Vec3{1, 0, 0}).Mul(dist)
	rot := mgl32.HomogRotate3DY(theta)
	res := rot.Mul4x1(disp.Vec4(1)).Vec3()

	return from.Add(res), at.Add(res)
}

func moveY(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	d := from.Sub(at)
	theta := float32(math.Atan(float64(d.Y() / d.Z())))
	disp := (mgl32.Vec3{0, 1, 0}).Mul(dist)
	rot := mgl32.HomogRotate3DX(theta)
	res := rot.Mul4x1(disp.Vec4(1)).Vec3()

	return from.Add(res), at.Add(res)
}

func rotX(from, at mgl32.Vec3, rad float32) (mgl32.Vec3, mgl32.Vec3) {
	d := from.Sub(at)
	rotY := mgl32.HomogRotate3DY(rad)
	return rotY.Mul4x1(d.Vec4(1)).Vec3().Add(at), at
}

func rotY(from, at mgl32.Vec3, rad float32) (mgl32.Vec3, mgl32.Vec3) {
	d := from.Sub(at)
	rot := mgl32.HomogRotate3DX(rad)
	return rot.Mul4x1(d.Vec4(1)).Vec3().Add(at), at
}

func cmdHandler(ch chan<- func(a, b mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3)) func(http.ResponseWriter, *http.Request) {
	cmdHandlers := make(map[string]func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3))
	cmdHandlers["zoom-in"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return zoom(from, at, 0.5)
	}

	cmdHandlers["zoom-out"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return zoom(from, at, -0.5)
	}

	cmdHandlers["left"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return moveX(from, at, -1)
	}

	cmdHandlers["right"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return moveX(from, at, 1)
	}

	cmdHandlers["up"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return moveY(from, at, 1)
	}

	cmdHandlers["down"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return moveY(from, at, -1)
	}
	cmdHandlers["rot-left"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return rotX(from, at, mgl32.DegToRad(-10))
	}

	cmdHandlers["rot-right"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return rotX(from, at, mgl32.DegToRad(10))
	}

	cmdHandlers["rot-up"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return rotY(from, at, mgl32.DegToRad(-10))
	}

	cmdHandlers["rot-down"] = func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
		return rotY(from, at, mgl32.DegToRad(10))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Path[len("/cmd/"):]

		if cmd, ok := cmdHandlers[cmd]; ok {
			ch <- cmd
		}

		fmt.Fprintf(w, cmd)
	}
}

func Start(ch chan<- func(a, b mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3)) error {
	http.Handle("/", http.FileServer(http.Dir("pkg/web/")))
	http.HandleFunc("/cmd/", cmdHandler(ch))
	return http.ListenAndServe(":8080", nil)
}
