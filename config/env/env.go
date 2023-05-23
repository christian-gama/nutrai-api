package env

// db is the database environment variables.
type db struct {
	// Host is the database host (IP address or domain) to connect to.
	Host DBHost `env:"DB_HOST,required"`

	// Port is the database port.
	Port DBPort `env:"DB_PORT,required"`

	// User is used in connection with the database password. This variable will create the
	// specified user with superuser power and a database with the same name.
	User DBUser `env:"DB_USER,required"`

	// Password must not be empty or undefined. This environment variable sets the superuser
	// password for
	// the database.
	Password DBPassword `env:"DB_PASSWORD,required"`

	// Name can be used to define a different name for the default database that is created when the
	// instance
	// is first started.
	Name DBName `env:"DB_NAME,required"`

	// SslMode is the database SSL mode. It expects "disable", "allow", "prefer", "require",
	// "verify-ca" or "verify-full". Please refer to the SQL driver documentation for more
	// information.
	SslMode DBSslMode `env:"DB_SSL_MODE,required"`
}

// jwt is the JWT environment variables.
type jwt struct {
	// Secret is the JWT secret key.
	Secret JwtSecret `env:"JWT_SECRET,required"`

	// AccessExpire is the JWT access token expiration time.
	AccessExpire JwtExpire `env:"JWT_ACCESS_EXPIRE,required"`

	// RefreshExpire is the JWT refresh token expiration time.
	RefreshExpire JwtExpire `env:"JWT_REFRESH_EXPIRE,required"`
}

// app is the application environment variables.
type app struct {
	// Host is the application host. It's usually the IP address of the machine
	Host AppHost `env:"APP_HOST,required"`

	// Port is the application port. It will be used to listen to incoming.
	Port AppPort `env:"APP_PORT,required"`

	// Env is the application environment. It expects "dev", "prod" or "test".
	Env AppEnv `env:"APP_ENV,required"`
}

// config is the configuration environment variables.
type config struct {
	// GlobalRateLimit is the global rate limit for the API. The expected value is in
	// requests per minute.
	GlobalRateLimit ConfigGlobalRateLimit `env:"CONFIG_GLOBAL_RATE_LIMIT,required"`

	// Debug is the debug mode. It will enable some debug features like the stack
	// trace in the response body.
	Debug bool `env:"CONFIG_DEBUG,required"`

	// LogLevel is the logging configuration for the application.
	LogLevel ConfigLogLevel `env:"CONFIG_LOG_LEVEL,required"`
}

type rabbitMQ struct {
	User     RabbitMQUser     `env:"RABBITMQ_USER,required"`
	Password RabbitMQPassword `env:"RABBITMQ_PASSWORD,required"`
	Host     RabbitMQHost     `env:"RABBITMQ_HOST,required"`
	Port     RabbitMQPort     `env:"RABBITMQ_PORT,required"`
}

// DB is the database environment variables.c.
var DB = &db{}

// App is the application environment variables.
var App = &app{}

// Config is the configuration environment variables.
var Config = &config{}

// RabbitMQ is the RabbitMQ environment variables.
var RabbitMQ = &rabbitMQ{}

// Jwt is the JWT environment variables.
var Jwt = &jwt{}

var (
	IsProduction  bool
	IsDevelopment bool
	IsTest        bool
)
