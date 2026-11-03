package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func Include(
	handler http.HandlerFunc,
	middlewares ...Middleware,
) http.Handler {
	return CreateStack(middlewares...)(handler)
}

/* Middleware Groups */
func Protected(handler http.Handler) http.Handler {
	return CreateStack(
		SetJsonResponseHeader,
	)(handler)
}

func Unprotected(handler http.Handler) http.Handler {
	return CreateStack(
		SetVaryAcceptResponseHeader,
	)(handler)
}

/* Middleware */
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func SetJsonResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func SetVaryAcceptResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept")

		next.ServeHTTP(w, r)
	})
}
