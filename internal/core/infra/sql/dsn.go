package sql

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/config/env"
)

// Dsn returns the database connection string.
func Dsn() string {
	host := fmt.Sprintf("host=%s", env.DB.Host)
	port := fmt.Sprintf("port=%d", env.DB.Port)
	dbname := fmt.Sprintf("dbname=%s", env.DB.Name)
	user := fmt.Sprintf("user=%s", env.DB.User)
	password := fmt.Sprintf("password=%s", env.DB.Password)
	sslmode := fmt.Sprintf("sslmode=%s", env.DB.SslMode)

	return fmt.Sprintf("%s %s %s %s %s %s", host, port, dbname, user, password, sslmode)
}
