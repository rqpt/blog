package router

import (
	"net/http"

	"github.com/rqpt/blog/internal/handlers"
	"github.com/rqpt/blog/internal/middleware"
)

type Handlers struct {
	Posts *handlers.PostHandler
}

func NewRouter(h Handlers) http.Handler {
	globalMiddleware := middleware.CreateStack(
		middleware.Logging,
	)

	router := http.NewServeMux()

	publicMux := newPublicMux(h)
	router.Handle("/", middleware.Unprotected(
		publicMux,
	))

	adminMux := newAdminMux(h)
	router.Handle("/admin/", middleware.Protected(
		http.StripPrefix("/admin", adminMux),
	))

	return globalMiddleware(router)
}
