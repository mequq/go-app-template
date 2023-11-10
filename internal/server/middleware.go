package server

import "log/slog"

type GorillaMuxMiddleware struct {
	logger    *slog.Logger
	level     slog.Level
	jwkSecret string
}

type MiddlewareOpt func(*GorillaMuxMiddleware)

func NewGorilaMuxAuthMiddleware(opt ...MiddlewareOpt) *GorillaMuxMiddleware {
	m := &GorillaMuxMiddleware{
		logger:    slog.Default(),
		jwkSecret: "secret",
		level:     slog.LevelInfo,
	}
	for _, o := range opt {
		o(m)
	}
	return m
}

func WithLogger(logger *slog.Logger) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) {
		m.logger = logger
	}
}

func WithLevel(level slog.Level) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) {
		m.level = level
	}
}

func WithJwkSecret(jwkSecret string) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) {
		m.jwkSecret = jwkSecret
	}
}
