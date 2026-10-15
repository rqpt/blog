package handlers

import "net/http"

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.postIndex)
	mux.HandleFunc("GET /{post}", h.postShow)

	mux.HandleFunc("POST /", h.postCreate)

	return mux
}
