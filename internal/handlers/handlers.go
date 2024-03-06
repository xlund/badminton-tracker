package handlers

import "net/http"

func LoadHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /", Home())
}
