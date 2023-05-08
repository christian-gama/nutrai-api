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

	validate()
}

func validate() {
	validSslModes := []string{"disable", "allow", "prefer", "require", "verify-ca", "verify-full"}
	if !slice.Contains(validSslModes, DB.SslMode) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be one of: %v",
				DB.SslMode,
				validSslModes,
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
		} else {
			os.Setenv("DB_HOST", "psql")
		}
	}
}
