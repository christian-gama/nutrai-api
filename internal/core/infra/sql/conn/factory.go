package conn

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
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
	conn := NewConn(postgres.Open, &gorm.Config{Logger: gormLogger}, log.MakeWithCaller())
	postgresConnection = conn.Open()

	return postgresConnection
}
