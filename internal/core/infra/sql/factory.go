package sql

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresConnection *gorm.DB

func MakePostgres() *gorm.DB {
	if postgresConnection != nil {
		return postgresConnection
	}

	conn := NewConn(postgres.Open, &gorm.Config{}, log.WithCaller)
	postgresConnection = conn.Open()

	return postgresConnection
}
