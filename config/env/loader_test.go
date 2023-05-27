package env_test

import (
	"fmt"
	"os"
	gopath "path"
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/pkg/path"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	testify "github.com/stretchr/testify/suite"
)

type LoaderSuite struct {
	suite.Suite
}

func TestLoaderSuite(t *testing.T) {
	testify.Run(t, new(LoaderSuite))
}

func (s *LoaderSuite) TestLoad() {
	s.Run("loads the environment variables", func() {
		s.NotPanics(func() {
			os.Clearenv()
			file, envFileName := createTempEnv(validEnvContent)
			defer os.Remove(file.Name())

			env.NewLoader(envFileName).Load()

			s.NotZero(env.DB.Host, "env.DB.Host")
			s.NotZero(env.DB.Name, "env.DB.Name")
			s.NotZero(env.DB.Password, "env.DB.Password")
			s.NotZero(env.DB.Port, "env.DB.Port")
			s.NotZero(env.DB.User, "env.DB.User")
			s.NotZero(env.DB.SslMode, "env.DB.SslMode")
		})

		s.Panics(func() {
			env.NewLoader("invalid").Load()
		})
	})
}

func createTempEnv(content string) (*os.File, string) {
	randomStr := fmt.Sprintf("%d", time.Now().UnixNano())
	envFileName := fmt.Sprintf(".env.temp.%s", randomStr)
	rootDir := path.Root()

	file, err := os.Create(gopath.Join(rootDir, envFileName))
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		os.Remove(file.Name())
		panic(err)
	}

	return file, envFileName
}

const validEnvContent = `
APP_HOST=localhost
APP_PORT=123
APP_ENV=testing
DB_HOST=localhost
DB_NAME=test
DB_PASSWORD=123
DB_PORT=123
DB_SSL_MODE=disable
DB_USER=123
DB_MAX_OPEN_CONNS=123
DB_MAX_IDLE_CONNS=123
DB_CONN_MAX_LIFETIME=1h
REDIS_HOST=localhost
REDIS_PORT=123
REDIS_PASSWORD=123
REDIS_USERNAME=123
CONFIG_GLOBAL_RATE_LIMIT=123
CONFIG_DEBUG=true
CONFIG_LOG_LEVEL=panic
JWT_SECRET=123
JWT_ACCESS_EXPIRE=123m
JWT_REFRESH_EXPIRE=123m
RABBITMQ_USER=123
RABBITMQ_PASSWORD=123
RABBITMQ_HOST=localhost
RABBITMQ_PORT=123
MAILER_PROVIDER=mailtrap
MAILER_FROM=from@email.com
MAILER_FROM_NAME=from_name
MAILTRAP_HOST=host
MAILTRAP_PORT=123
MAILTRAP_USER=user
MAILTRAP_PASSWORD=password
SENDGRID_API_KEY=123
`
