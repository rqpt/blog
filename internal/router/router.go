package router

import (
	"net/http"

	"github.com/rqpt/blog/internal/handlers"
)

type Handlers struct {
	Posts *handlers.PostHandler
}

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()

	// TODO: post create, update and delete should have some kind of auth
	// middleware, and request rate limit associated with them.
	mux.HandleFunc("POST /", h.Posts.Create)
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
