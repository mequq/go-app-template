package server

import (
	"application/internal/utils"
	"net/http"
)

// NewGorillaMuxServer creates a new HTTP server and set up all routes.
func (m *GorillaMuxMiddleware) ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = utils.SetContextFromHttpReq(ctx, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
