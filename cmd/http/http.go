package http

import (
	"context"

	"github.com/bangcodv/cake-service/internal/consts"
	"github.com/bangcodv/cake-service/internal/server"
	"github.com/bangcodv/cake-service/pkg/logger"
)

// Start function handler starting http listener
func Start(ctx context.Context) {

	serve := server.NewHTTPServer()
	defer serve.Done()
	logger.Info(logger.MessageFormat("starting cake-service services... %d", serve.Config().App.Port), logger.EventName(consts.LogEventNameServiceStarting))

	if err := serve.Run(ctx); err != nil {
		logger.Warn(logger.MessageFormat("service stopped, err:%s", err.Error()), logger.EventName(consts.LogEventNameServiceStarting))
	}

	return
}
