package router

import "net/http"

// TODO: admin routes should have some kind of auth
// middleware, and request rate limit associated with them.
func newAdminMux(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.Posts.Create)
	mux.HandleFunc("PUT /{post}", h.Posts.Update)
	mux.HandleFunc("DELETE /{post}", h.Posts.Delete)

	return mux
}
