package env

import "os"

// docker represents an environment running within a Docker container.
type docker struct {
	envFile string
}

// newDocker initializes a new Docker instance.
// It takes the name of the .env file as a parameter.
func newDocker(envFile string) *docker {
	return &docker{
		envFile: envFile,
	}
}

// check verifies if the application is running inside a Docker container.
// If it is, it sets up the environment variables for the DB host and RabbitMQ host.
func (d *docker) check() {
	runningInDocker := os.Getenv("RUNNING_IN_DOCKER")
	if runningInDocker != "true" {
		return
	}

	d.setupDBHost()
	d.setupRabbitMQHost()
}

// setupDBHost sets up the DB host environment variable.
// The host value changes depending on whether the application is running in the test environment.
func (d *docker) setupDBHost() {
	dbHost := "psql"

	if d.isTestEnv() {
		dbHost = "psql_test"
	}

	os.Setenv("DB_HOST", dbHost)
}

// setupRabbitMQHost sets up the RabbitMQ host environment variable.
// The host value changes depending on whether the application is running in the test environment.
func (d *docker) setupRabbitMQHost() {
	rabbitMQHost := "rabbitmq"

	if d.isTestEnv() {
		rabbitMQHost = "rabbitmq_test"
	}

	os.Setenv("RABBITMQ_HOST", rabbitMQHost)
}

// isTestEnv checks if the application is running in the test environment.
func (d *docker) isTestEnv() bool {
	return d.envFile == ".env.test"
}
