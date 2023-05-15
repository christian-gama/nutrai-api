package env

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
