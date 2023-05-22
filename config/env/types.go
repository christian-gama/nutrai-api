package env

import (
	"strconv"
	"strings"
	"time"
)

// Application.
type (
	AppEnv  string
	AppPort int
	AppHost string
)

// Database.
type (
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSslMode  string
)

// Config.
type (
	ConfigGlobalRateLimit int
	ConfigDebug           bool
)

// RabbitMQ.
type (
	RabbitMQHost     string
	RabbitMQPort     int
	RabbitMQUser     string
	RabbitMQPassword string
)

// JWT.
type (
	JwtSecret string
	JwtExpire struct{ time.Duration }
)

// UnmarshalText parses a time duration from text, including support for "d" suffix.
func (d *JwtExpire) UnmarshalText(text []byte) error {
	s := string(text)
	if strings.HasSuffix(s, "d") {
		days, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return err
		}
		d.Duration = time.Duration(days) * 24 * time.Hour
	} else {
		var err error
		d.Duration, err = time.ParseDuration(s)
		if err != nil {
			return err
		}
	}
	return nil
}
