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
	theta := float32(math.Atan(float64(d.Z() / d.X())))
	rotT := mgl32.HomogRotate3DY(-theta)
	rotP := mgl32.HomogRotate3DY(theta)
	rotY := mgl32.HomogRotate3DY(rad)
	rot := rotP.Mul4(rotY).Mul4(rotT)
	return rot.Mul4x1(d.Vec4(1)).Vec3().Add(at), at
}

func rotY(from, at mgl32.Vec3, rad float32) (mgl32.Vec3, mgl32.Vec3) {
	d := from.Sub(at)
	theta := float32(math.Atan(float64(d.Z() / d.Y())))
	fmt.Println(theta)
	rotT := mgl32.HomogRotate3DX(-theta)
	rotP := mgl32.HomogRotate3DX(theta)
	rotZ := mgl32.HomogRotate3DX(rad)
	rot := rotP.Mul4(rotZ).Mul4(rotT)
	return rot.Mul4x1(d.Vec4(1)).Vec3().Add(at), at
}

func cmdHandler(ch chan<- func(a, b mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Path[len("/cmd/"):]

		switch cmd {
		case "zoom-in":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return zoom(from, at, 0.5)
			}

		case "zoom-out":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return zoom(from, at, -0.5)
			}

		case "left":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveX(from, at, -1)
			}

		case "right":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveX(from, at, 1)
			}

		case "up":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveY(from, at, 1)
			}

		case "down":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveY(from, at, -1)
			}
		case "rot-left":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return rotX(from, at, mgl32.DegToRad(-10))
			}

		case "rot-right":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return rotX(from, at, mgl32.DegToRad(10))
			}

		case "rot-up":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return rotY(from, at, mgl32.DegToRad(-10))
			}

		case "rot-down":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return rotY(from, at, mgl32.DegToRad(10))
			}

		}

		fmt.Fprintf(w, cmd)
	}
}

func Start(ch chan<- func(a, b mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3)) {
	http.Handle("/", http.FileServer(http.Dir("pkg/web/")))

	http.HandleFunc("/cmd/", cmdHandler(ch))

	http.ListenAndServe(":8080", nil)
}

func Close() {
}
