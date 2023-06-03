package conn

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MakePsql() {
	if psql != nil {
		return
	}

	gormLogger := logger.Discard
	if env.Config.Debug {
		gormLogger = logger.Default
	}

	psql = NewConn(postgres.Open, &gorm.Config{Logger: gormLogger})
}
