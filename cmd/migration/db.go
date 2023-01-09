// Package migration
package migration

import (
	"time"

	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/pkg/databasex"
)

func MigrateDatabase() {
	cfg := appctx.NewConfig()

	databasex.DatabaseMigration(&databasex.Config{
		Driver:       cfg.WriteDB.Driver,
		Host:         cfg.WriteDB.Host,
		Port:         cfg.WriteDB.Port,
		Name:         cfg.WriteDB.Name,
		User:         cfg.WriteDB.User,
		Password:     cfg.WriteDB.Pass,
		Charset:      cfg.WriteDB.Charset,
		Timeout:      cfg.WriteDB.TimeoutSecond,
		MaxIdleConns: cfg.WriteDB.MaxIdle,
		MaxOpenConns: cfg.WriteDB.MaxOpen,
		MaxLifetime:  time.Duration(cfg.WriteDB.MaxLifeTimeMS) * time.Millisecond,
		TimeZone:     cfg.WriteDB.Timezone,
	})
}
