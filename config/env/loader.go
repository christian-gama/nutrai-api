package env

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/christian-gama/nutrai-api/pkg/path"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type loader struct {
	ctx     context.Context
	envFile string
	docker  *docker
}

// NewLoader returns a new instance of loader. It receives the name of the environment file
// (envFile) to be loaded.
func NewLoader(envFile string) *loader {
	return &loader{
		ctx:     context.Background(),
		envFile: envFile,
		docker:  newDocker(envFile),
	}
}

// Load loads the environment variables from the given environment file (envFile) and validates
// them.
func (e *loader) Load() {
	e.docker.check()

	err := godotenv.Load(Path(e.envFile))
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %w", err))
	}

	e.loadEnvironmentVariables()
	e.setEnvironmentType()

	e.validate(
		validateAppEnv(),
		validateConfigLogLevel(),
		validateDBSslMode(),
		validateMailerProvider(),
	)
}

// loadEnvironmentVariables loads the environment variables into the given variables.
func (e *loader) loadEnvironmentVariables() {
	variables := []any{DB, App, Config, Jwt, RabbitMQ, Mailer, Mailtrap, Sendgrid, Redis, Gpt}

	for _, variable := range variables {
		if err := envconfig.Process(e.ctx, variable); err != nil {
			panic(fmt.Errorf("Error loading environment variables: %w", err))
		}
	}
}

func (e *loader) setEnvironmentType() {
	IsProduction = App.Env == Production
	IsDevelopment = App.Env == Development
	IsTest = App.Env == Test
}

func (e *loader) validate(validators ...validator) {
	for _, validator := range validators {
		if err := validator(); err != nil {
			panic(err)
		}
	}
}

// Path returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func Path(envFile string) string {
	rootDir := path.Root()
	return filepath.Join(rootDir, envFile)
}
