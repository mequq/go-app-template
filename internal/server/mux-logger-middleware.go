package server

import (
	"application/internal/utils"
	"log/slog"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (m *GorillaMuxMiddleware) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recored time duration
		startTime := time.Now()

		ctx := r.Context()
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(recorder, r)
		m.logger.Log(ctx, m.level, "request fulfilled",
			slog.Group(
				"request-info",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
			),
			"ctx", utils.LogContext(ctx),
			"status", recorder.Status,
			"duration", time.Since(startTime).String(),
		)
	})
}
