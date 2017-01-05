package web

import (
	"fmt"
	"net/http"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Path[len("/cmd/"):]
	fmt.Fprintf(w, cmd)
}

func Start() {
	http.HandleFunc("/cmd/", cmdHandler)

	http.ListenAndServe(":8080", nil)
}

func Close() {
}
