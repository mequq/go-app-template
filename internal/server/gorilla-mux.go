package server

import (
	"application/config"
	"application/internal/service"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// New GorilaMuxServer creates a new HTTP server and set up all routes.

func NewGorillaMuxServer(
	cfg *config.ViperConfig,
	logger *slog.Logger,
	healthzSvc *service.GorilaMuxHealthzService,

) http.Handler {

	mux := mux.NewRouter()
	middleware := NewGorilaMuxAuthMiddleware(WithLevel(slog.LevelInfo), WithLogger(logger))
	// logger middleware

	mux.Use(otelmux.Middleware("my-server"))
	mux.Use(middleware.ContextMiddleware)
	mux.Use(middleware.RecoverMiddleware)
	mux.Use(middleware.LoggerMiddleware)

	healthzSvc.RegisterRoutes(mux)
	// http.Handle("/", mux)
	return mux

}
