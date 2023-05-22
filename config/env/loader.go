package env

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/christian-gama/nutrai-api/pkg/path"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// Load loads the environment variables from the .env file.
func Load(envFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	checkDocker(envFile)

	err := godotenv.Load(Path(envFile))
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %w", err))
	}

	if err := envconfig.Process(ctx, DB); err != nil {
		panic(fmt.Errorf("Error loading DB environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, App); err != nil {
		panic(fmt.Errorf("Error loading App environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, Config); err != nil {
		panic(fmt.Errorf("Error loading Config environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, Jwt); err != nil {
		panic(fmt.Errorf("Error loading Jwt environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, RabbitMQ); err != nil {
		panic(fmt.Errorf("Error loading RabbitMQ environment variables: %w", err))
	}

	validate()
}

func validate() {
	validSslModes := []DBSslMode{
		SslModeDisable,
		SslModeAllow,
		SslModePrefer,
		SslModeRequire,
		SslModeVerifyCa,
		SslModeVerifyFull,
	}
	if !slice.Contains(validSslModes, DB.SslMode) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be one of: %v",
				DB.SslMode,
				validSslModes,
			),
		)
	}

	validEnvs := []AppEnv{Production, Development, Test}
	if !slice.Contains(validEnvs, App.Env) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be one of: %v",
				App.Env,
				validEnvs,
			),
		)
	}

	validLogLevels := []ConfigLogLevel{
		LogLevelInfo,
		LogLevelWarn,
		LogLevelError,
		LogLevelDebug,
		LogLevelPanic,
	}
	if !slice.Contains(validLogLevels, Config.LogLevel) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be one of: %v",
				Config.LogLevel,
				validLogLevels,
			),
		)
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

func checkDocker(envFile string) {
	runningInDocker := os.Getenv("RUNNING_IN_DOCKER")

	if runningInDocker == "true" {
		if envFile == ".env.test" {
			os.Setenv("DB_HOST", "psql_test")
			os.Setenv("RABBITMQ_HOST", "rabbitmq_test")
		} else {
			os.Setenv("DB_HOST", "psql")
			os.Setenv("RABBITMQ_HOST", "rabbitmq")
		}
	}
}
