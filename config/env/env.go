package env

// db is the database environment variables.
type db struct {
	// Host is the database host (IP address or domain) to connect to.
	Host string `env:"DB_HOST,required"`

	// Port is the database port.
	Port int `env:"DB_PORT,required"`

	// User is used in connection with the database password. This variable will create the
	// specified user with superuser power and a database with the same name.
	User string `env:"DB_USER,required"`

	// Password must not be empty or undefined. This environment variable sets the superuser
	// password for	the database.
	Password string `env:"DB_PASSWORD,required"`

	// Name can be used to define a different name for the default database that is created when the
	// instance is first started.
	Name string `env:"DB_NAME,required"`

	// SslMode is the database SSL mode. It expects "disable", "allow", "prefer", "require",
	// "verify-ca" or "verify-full". Please refer to the SQL driver documentation for more
	// information.
	SslMode DBSslMode `env:"DB_SSL_MODE,required"`
}

// jwt is the JWT environment variables.
type jwt struct {
	// Secret is the JWT secret key.
	Secret string `env:"JWT_SECRET,required"`

	// AccessExpire is the JWT access token expiration time.
	AccessExpire JwtExpire `env:"JWT_ACCESS_EXPIRE,required"`

	// RefreshExpire is the JWT refresh token expiration time.
	RefreshExpire JwtExpire `env:"JWT_REFRESH_EXPIRE,required"`
}

// app is the application environment variables.
type app struct {
	// Host is the application host. It's usually the IP address of the machine
	Host string `env:"APP_HOST,required"`

	// Port is the application port. It will be used to listen to incoming.
	Port int `env:"APP_PORT,required"`

	// Env is the application environment. It expects "dev", "prod" or "test".
	Env AppEnv `env:"APP_ENV,required"`
}

// config is the configuration environment variables.
type config struct {
	// GlobalRateLimit is the global rate limit for the API. The expected value is in
	// requests per minute.
	GlobalRateLimit int `env:"CONFIG_GLOBAL_RATE_LIMIT,required"`

	// Debug is the debug mode. It will enable some debug features like the stack
	// trace in the response body.
	Debug bool `env:"CONFIG_DEBUG,required"`

	// LogLevel is the logging configuration for the application.
	LogLevel ConfigLogLevel `env:"CONFIG_LOG_LEVEL,required"`
}

type rabbitMQ struct {
	// User is the RabbitMQ user.User name to create when RabbitMQ creates a new database from
	// scratch.
	User string `env:"RABBITMQ_USER,required"`

	// Password is the default user password.
	Password string `env:"RABBITMQ_PASSWORD,required"`

	// Host is the RabbitMQ host (IP address or domain) to connect to.
	Host string `env:"RABBITMQ_HOST,required"`

	// RabbitMQ nodes will use a port from a certain range known as the inter-node communication
	// port range. The same port is used by CLI tools when they need to contact the node. The range
	// can be modified.
	Port int `env:"RABBITMQ_PORT,required"`
}

type mailer struct {
	// Provider is the mailer provider. It expects "mailtrap", "sendgrid" or "mailgun".
	Provider MailerProvider `env:"MAILER_PROVIDER,required"`

	// From is the mailer from address.
	From string `env:"MAILER_FROM,required"`

	// FromName is the mailtrap from name.
	FromName string `env:"MAILER_FROM_NAME,required"`
}

type mailtrap struct {
	// Host is the mailtrap host (IP address or domain) to connect to.
	Host string `env:"MAILTRAP_HOST,required"`

	// Port is the mailtrap port.
	Port int `env:"MAILTRAP_PORT,required"`

	// Username is the mailtrap user.
	Username string `env:"MAILTRAP_USERNAME,required"`

	// Password is the mailtrap password.
	Password string `env:"MAILTRAP_PASSWORD,required"`
}

type sendgrid struct {
	// ApiKey is the sendgrid API key.
	ApiKey string `env:"SENDGRID_API_KEY,required"`
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

// Mailer is the mailer environment variables.
var Mailer = &mailer{}

// Mailtrap is the mailtrap environment variables.
var Mailtrap = &mailtrap{}

// Sendgrid is the sendgrid environment variables.
var Sendgrid = &sendgrid{}

var (
	// IsProduction is true if the application is running in production mode.
	IsProduction bool

	// IsDevelopment is true if the application is running in development mode.
	IsDevelopment bool

	// IsTest is true if the application is running in test mode.
	IsTest bool
)
