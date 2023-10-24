package server

import "log/slog"

type GorilaMuxAuthMiddleware struct {
	logger    *slog.Logger
	authToken string
}

func NewGorilaMuxAuthMiddleware(logger *slog.Logger, authToken string, opt ...func(*GorilaMuxAuthMiddleware)) *GorilaMuxAuthMiddleware {
	m := &GorilaMuxAuthMiddleware{
		logger:    logger.With("layer", "GorilaMuxAuthMiddleware"),
		authToken: authToken,
	}
	for _, o := range opt {
		o(m)
	}
	return m
}
