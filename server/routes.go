package server

import "net/http"

func addRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("/", http.FileServer(http.Dir("static")))
}
