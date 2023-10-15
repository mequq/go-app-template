package service

import (
	"application/internal/biz"
	"net/http"

	"log/slog"

	amid "git.abanppc.com/lenz-public/gin-utils/middleware"
	"github.com/gin-gonic/gin"
)

type HealthzService struct {
	uc     biz.HealthzUseCaseInterface
	logger *slog.Logger
}

// New
func NewHealthzService(uc biz.HealthzUseCaseInterface, logger *slog.Logger) *HealthzService {
	return &HealthzService{
		uc:     uc,
		logger: logger.With("layer", "HealthzService"),
	}
}

// Healthz Liveness
func (s *HealthzService) HealthzLiveness(c *gin.Context) {
	ctx := amid.GetContextFromGin(c)
	s.logger.DebugContext(ctx, "HealthzLiveness", "ctx", c.Request.URL)
	s.uc.Liveness(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// Healthz Readiness
func (s *HealthzService) HealthzReadiness(c *gin.Context) {
	ctx := amid.GetContextFromGin(c)
	s.logger.DebugContext(ctx, "HealthzReadiness", "ctx", ctx)
	s.uc.Readiness(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// Healthz Route
func (s *HealthzService) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/healthz")
	r.GET("/liveness", s.HealthzLiveness)
	r.GET("/readiness", s.HealthzReadiness)
}
