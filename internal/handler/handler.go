package handler

import "net/http"

type Handler struct {}

func(h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /weather/{city}", h.weather)

	return mux
}