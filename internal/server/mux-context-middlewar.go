package server

import (
	"application/internal/biz"
	"net/http"
)

// NewGorillaMuxServer creates a new HTTP server and set up all routes.

func GorillaMuxContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = biz.SetContextFromHttpReq(ctx, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
