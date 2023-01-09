// Package bootstrap
package bootstrap

import (
	"github.com/bangcodv/cake-service/internal/consts"
	"github.com/bangcodv/cake-service/pkg/logger"
	"github.com/bangcodv/cake-service/pkg/msgx"
)

func RegistryMessage()  {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
