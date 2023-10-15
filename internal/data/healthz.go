package data

import (
	"application/internal/biz"
	"context"
	"log/slog"
)

type HealthzRepo struct {
	logger *slog.Logger
	ds     *DataSource
}

func NewHealthzRepo(logger *slog.Logger, ds *DataSource) biz.HealthzRepoInterface {
	return &HealthzRepo{
		logger: logger.With("repo", "HealthzRepo"),
		ds:     ds,
	}
}

func (r *HealthzRepo) Readiness(ctx context.Context) error {
	r.logger.Debug("repo Readiness")
	return nil
}

func (r *HealthzRepo) Liveness(ctx context.Context) error {
	r.logger.Debug("repo Liveness")
	return nil
}
