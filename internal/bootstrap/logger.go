// Package bootstrap
package bootstrap

import (
	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/pkg/logger"
	"github.com/bangcodv/cake-service/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}

