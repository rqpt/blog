package router

import "net/http"

// TODO: admin routes should have some kind of auth
// middleware, and request rate limit associated with them.
func newAdminMux(h Handlers) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", h.Posts.Create)
	router.HandleFunc("PUT /{post}", h.Posts.Update)
	router.HandleFunc("DELETE /{post}", h.Posts.Delete)

	return router
}
