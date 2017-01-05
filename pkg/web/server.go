package web

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"net/http"
)

func zoom(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	displacement := from.Sub(at).Normalize()
	return from.Sub(displacement.Mul(dist)), at
}

func moveX(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	return from.Sub(mgl32.Vec3{dist, 0, 0}), at.Sub(mgl32.Vec3{dist, 0, 0})
}

func moveY(from, at mgl32.Vec3, dist float32) (mgl32.Vec3, mgl32.Vec3) {
	return from.Sub(mgl32.Vec3{0, dist, 0}), at.Sub(mgl32.Vec3{0, dist, 0})
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
				return moveX(from, at, 1)
			}

		case "right":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveX(from, at, -1)
			}

		case "up":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveY(from, at, -1)
			}

		case "down":
			ch <- func(from, at mgl32.Vec3) (mgl32.Vec3, mgl32.Vec3) {
				return moveY(from, at, 1)
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
