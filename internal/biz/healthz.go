package biz

import (
	"context"

	"log/slog"
)

type HealthzRepoInterface interface {
	Readiness(ctx context.Context) error
	Liveness(ctx context.Context) error
}

type HealthzUseCaseInterface interface {
	Readiness(ctx context.Context) error
	Liveness(ctx context.Context) error
}

type HealthzUseCase struct {
	repo   HealthzRepoInterface
	logger *slog.Logger
}

// New Usecase
func NewHealthzUseCase(repo HealthzRepoInterface, logger *slog.Logger) HealthzUseCaseInterface {
	return &HealthzUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *HealthzUseCase) Readiness(ctx context.Context) error {
	return uc.repo.Readiness(ctx)
}

func (uc *HealthzUseCase) Liveness(ctx context.Context) error {
	return uc.repo.Liveness(ctx)
}
