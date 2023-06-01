package conn

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var postgresConnection *gorm.DB

func MakePostgres() *gorm.DB {
	if postgresConnection != nil {
		return postgresConnection
	}

	gormLogger := logger.Discard
	if env.Config.Debug {
		gormLogger = logger.Default
	}

	postgresConnection = NewConn(postgres.Open, &gorm.Config{Logger: gormLogger})

	return postgresConnection
}
