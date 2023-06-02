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
APP_ENV=test
APP_HOST=localhost
APP_PORT=8080
APP_ALLOWED_ORIGINS=*
CONFIG_GLOBAL_RATE_LIMIT=0
CONFIG_ENABLE_RATE_LIMIT=false
CONFIG_DEBUG=false
CONFIG_LOG_LEVEL=panic
DB_HOST=localhost
DB_NAME=db-name
DB_PASSWORD=password
DB_PORT=5434
DB_SSL_MODE=disable
DB_USER=username
DB_MAX_OPEN_CONNS=100
DB_MAX_IDLE_CONNS=10
DB_CONN_MAX_LIFETIME=1h
REDIS_HOST=localhost
REDIS_PORT=6381
REDIS_PASSWORD=password
JWT_SECRET=secret
JWT_ACCESS_EXPIRE=999d
JWT_REFRESH_EXPIRE=999d
JWT_AUDIENCE=https://yourdomain.com
JWT_ISSUER=https://yourdomain.com
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5674
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
MAILER_PROVIDER=mailtrap
MAILER_TEMPLATE_PATH=templates
MAILER_ASSETS_PATH=templates/assets
MAILER_FROM=nutrai.team@gmail.com
MAILER_FROM_NAME=Nutrai Team
MAILER_DELAY_BETWEEN_EMAILS=3s
MAILTRAP_HOST=sandbox.smtp.mailtrap.io
MAILTRAP_PORT=2525
MAILTRAP_USERNAME=your_mailtrap_user
MAILTRAP_PASSWORD=your_mailtrap_password
SENDGRID_API_KEY=your_sendgrid_api_key
`
