package routes

import (
	"io"
	"net/http"
)

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "PONG")
	})

	return mux
}
