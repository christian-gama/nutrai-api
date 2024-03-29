package env

import "time"

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

	// MaxIdleConns is the maximum number of connections in the idle connection pool.
	MaxIdleConns int `env:"DB_MAX_IDLE_CONNS,required"`

	// MaxOpenConns is the maximum number of open connections to the database.
	MaxOpenConns int `env:"DB_MAX_OPEN_CONNS,required"`

	// ConnMaxLifetime is the maximum amount of time (in minutes) a connection may be reused.
	ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME,required"`
}

// redis is the Redis environment variables.
type redis struct {
	// Host is the Redis host (IP address or domain) to connect to.
	Host string `env:"REDIS_HOST,required"`

	// Port is the Redis port.
	Port int `env:"REDIS_PORT,required"`

	// Password is the Redis password.
	Password string `env:"REDIS_PASSWORD,required"`
}

// jwt is the JWT environment variables.
type jwt struct {
	// Secret is the JWT secret key.
	Secret string `env:"JWT_SECRET,required"`

	// AccessExpire is the JWT access token expiration time.
	AccessExpire JwtExpire `env:"JWT_ACCESS_EXPIRE,required"`

	// RefreshExpire is the JWT refresh token expiration time.
	RefreshExpire JwtExpire `env:"JWT_REFRESH_EXPIRE,required"`

	// Audience is the JWT audience.
	Audience string `env:"JWT_AUDIENCE,required"`

	// Issuer is the JWT issuer.
	Issuer string `env:"JWT_ISSUER,required"`
}

// app is the application environment variables.
type app struct {
	// Host is the application host. It's usually the IP address of the machine
	Host string `env:"APP_HOST,required"`

	// Port is the application port. It will be used to listen to incoming.
	Port int `env:"APP_PORT,required"`

	// Env is the application environment. It expects "dev", "prod" or "test".
	Env AppEnv `env:"APP_ENV,required"`

	// AllowedOrigins is the list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	AllowedOrigins []string `env:"APP_ALLOWED_ORIGINS,required"`

	// Key is the application key. It's usually passed in the request header and it's used to
	// authenticate the request.
	ApiKey string `env:"APP_API_KEY,required"`
}

// config is the configuration environment variables.
type config struct {
	// GlobalRateLimit is the global rate limit for the API. The expected value is in
	// requests per minute.
	GlobalRateLimit int `env:"CONFIG_GLOBAL_RATE_LIMIT,required"`

	// EnableRateLimit is the flag that indicates if the rate limit is enabled or not.
	EnableRateLimit bool `env:"CONFIG_ENABLE_RATE_LIMIT,required"`

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

	// TemplatePath is the mailer template path.
	TemplatePath string `env:"MAILER_TEMPLATE_PATH,required"`

	// AssetsPath is the mailer template assets path.
	AssetsPath string `env:"MAILER_ASSETS_PATH,required"`

	// FromName is the mailtrap from name.
	FromName string `env:"MAILER_FROM_NAME,required"`

	// MAILER_DELAY_BETWEEN_EMAILS is the delay that will be used to send each email. It's used to
	// throttle the email sending process, to avoid being blocked by the email provider.
	DelayBetweenEmails time.Duration `env:"MAILER_DELAY_BETWEEN_EMAILS,required"`
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

type gpt struct {
	// Model is the GPT model.
	Model string `env:"GPT_DEFAULT_MODEL,required"`

	// MaxTokens is the GPT max tokens.
	MaxTokens int `env:"GPT_DEFAULT_MODEL_MAX_TOKENS,required"`

	// Stop is the GPT stop tokens.
	Stop []string `env:"GPT_DEFAULT_STOP,required"`

	// Temperature is the GPT temperature.
	Temperature float32 `env:"GPT_DEFAULT_TEMPERATURE,required"`

	// TopP is the GPT top p.
	TopP float32 `env:"GPT_DEFAULT_TOP_P,required"`

	// N is number of completions to generate for each prompt.
	N int `env:"GPT_DEFAULT_N,required"`

	// PresencePenalty is the GPT presence penalty.
	PresencePenalty float32 `env:"GPT_DEFAULT_PRESENCE_PENALTY"`

	// FrequencyPenalty is the GPT frequency penalty.
	FrequencyPenalty float32 `env:"GPT_DEFAULT_FREQUENCY_PENALTY"`

	// ApiKey is the GPT API key.
	ApiKey string `env:"GPT_API_KEY,required"`
}

// DB is the database environment variables.c.
var DB = &db{}

// Redis is the Redis environment variables.
var Redis = &redis{}

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

// Gpt is the gpt environment variables.
var Gpt = &gpt{}

var (
	// IsProduction is true if the application is running in production mode.
	IsProduction bool

	// IsDevelopment is true if the application is running in development mode.
	IsDevelopment bool

	// IsTest is true if the application is running in test mode.
	IsTest bool
)
