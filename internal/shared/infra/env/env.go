package env

// db is the database environment variables.
type db struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	SslMode  string `env:"DB_SSL_MODE,required"`
}

// DB is the database environment variables.c.
var DB = &db{}
