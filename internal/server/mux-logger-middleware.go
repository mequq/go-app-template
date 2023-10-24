package server

import (
	"application/internal/biz"
	"log/slog"
	"net/http"
	"time"
)

// import (
// 	"net/http"

// 	"log/slog"
// )

type GorillaMuxLoggerMiddleware struct {
	logger *slog.Logger
	level  slog.Level
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func NewGorillaMuxLoggerMiddleware(opt ...func(*GorillaMuxLoggerMiddleware)) GorillaMuxLoggerMiddleware {
	mid := GorillaMuxLoggerMiddleware{
		logger: slog.Default(),
		level:  slog.LevelInfo,
	}
	for _, o := range opt {
		o(&mid)
	}
	return mid
}

func WithLogger(logger *slog.Logger) func(*GorillaMuxLoggerMiddleware) {
	return func(m *GorillaMuxLoggerMiddleware) {
		m.logger = logger
	}
}

func WithLevel(level slog.Level) func(*GorillaMuxLoggerMiddleware) {
	return func(m *GorillaMuxLoggerMiddleware) {
		m.level = level
	}
}

func (m *GorillaMuxLoggerMiddleware) Middleware(next http.Handler) http.Handler {
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
			"ctx", biz.LogContext(ctx),
			"status", recorder.Status,
			"duration", time.Since(startTime).String(),
		)
	})
}
