//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"application/config"
	"application/internal/biz"
	"application/internal/data"
	"application/internal/server"
	"application/internal/service"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func wireApp(cfg *config.ViperConfig, logger *slog.Logger) (*gin.Engine, error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
	))
}
