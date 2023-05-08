package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresConnection *gorm.DB

func MakePostgres() *gorm.DB {
	if postgresConnection != nil {
		return postgresConnection
	}

	conn := NewConn(postgres.Open, &gorm.Config{})
	postgresConnection = conn.Open()

	return postgresConnection
}
