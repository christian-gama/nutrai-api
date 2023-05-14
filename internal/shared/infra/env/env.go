package env

// db is the database environment variables.
type db struct {
	Host     DBHost     `env:"DB_HOST,required"`
	Port     DBPort     `env:"DB_PORT,required"`
	User     DBUser     `env:"DB_USER,required"`
	Password DBPassword `env:"DB_PASSWORD,required"`
	Name     DBName     `env:"DB_NAME,required"`
	SslMode  DBSslMode  `env:"DB_SSL_MODE,required"`
}

// app is the application environment variables.
type app struct {
	Host AppHost `env:"APP_HOST,required"`
	Port AppPort `env:"APP_PORT,required"`
	Env  AppEnv  `env:"APP_ENV,required"`
}

// config is the configuration environment variables.
type config struct {
	GlobalRateLimit ConfigGlobalRateLimit `env:"CONFIG_GLOBAL_RATE_LIMIT,required"`
	Debug           ConfigDebug           `env:"CONFIG_DEBUG,required"`
}

// DB is the database environment variables.c.
var DB = &db{}

// App is the application environment variables.
var App = &app{}

// Config is the configuration environment variables.
var Config = &config{}
