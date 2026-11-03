package router

import "net/http"

func newPublicMux(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", h.Posts.Index)
	mux.HandleFunc("GET /{post}", h.Posts.Show)
	mux.Handle("GET /static/", static())

	return mux
}

func static() http.Handler {
	return http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir(
			"internal/renderer/resources/public",
		)),
	)
}
