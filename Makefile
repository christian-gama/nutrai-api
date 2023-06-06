# ==============================================================================================
# Title: Makefile
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-07
# Usage:    Run the command 'make help' to see all available commands.
# ==============================================================================================

# APP_NAME is the name of the application.
APP_NAME = nutrai-api

# AIRVERSION is the version of the air package to be used. 
# https://github.com/cosmtrek/air
AIRVERSION = v1.43.0

# WORKDIR is used to set the working directory for Dockerfile builds.
WORKDIR = /usr/src/app

# ==============================================================================================
# Target: help
# Brief: Displays a brief description of all the targets in the Makefile.
# Usage: Run the command 'make help'.
# ==============================================================================================
.PHONY: help
help: .clear-screen
	@sh -c "./scripts/help.sh"

# ==============================================================================================
# Target: init
# Brief: Initializes the project.
# Usage: Run the command 'make init'.
# ==============================================================================================
.PHONY: init
init: .cmd-exists-git .cmd-exists-go .cmd-exists-docker .cmd-exists-sh .clear-screen
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Git Hooks configured."

	@chmod +x ./scripts/*.sh
	@echo "Scripts configured."
	@sh ./scripts/manage_env.sh


# ==============================================================================================
# Target: env
# Brief: Creates or updates the environment file based on the .env.example file.
# Usage: Run the command 'make env'.
# ==============================================================================================
.PHONY: env
env: .cmd-exists-sh .clear-screen
	@sh ./scripts/manage_env.sh
	@$(MAKE) metric-key


# ==============================================================================================
# Target: metric-key
# Brief: Generates a new metric key.
# Usage: Run the command 'make metric-key'.
# Flags:
# 	ENV_FILE: The path to the environment file.
# =============================================================================================
.PHONY: metric-key .check-env-file .clear-screen
metric-key:
	@./scripts/metric_key.sh


# ==============================================================================================
# Target: run
# Brief: Runs the API server.
# Usage: Run the command 'make run [ENV_FILE="<path>"]'.
# Flags:
#  ENV_FILE: The path to the environment file. It defaults to '.env.dev'.
# ==============================================================================================
.PHONY: run
run: .cmd-exists-go .clear-screen .check-env-file
ifneq ($(RUNNING_IN_DOCKER), true)
	@$(MAKE) postgres
	@$(MAKE) rabbitmq
	@$(MAKE) redis
endif

ifeq ($(ENV_FILE), .env.prod)
	@$(MAKE) build
	@$(BUILD_DIR)/$(APP_NAME) -e $(ENV_FILE)
else
	@go run github.com/cosmtrek/air@$(AIRVERSION) -- -e $(ENV_FILE)
endif


# ==============================================================================================
# Target: dev
# Brief: Runs the API server in development mode.
# Usage: Run the command 'make dev'.
# ==============================================================================================
.PHONY: dev
dev: .cmd-exists-go .clear-screen
	@ENV_FILE=.env.dev $(MAKE) run


# ==============================================================================================
# Target: prod
# Brief: Runs the API server in production mode.
# Usage: Run the command 'make prod'.
# ==============================================================================================
.PHONY: prod
prod: .cmd-exists-go .clear-screen
	@ENV_FILE=.env.prod $(MAKE) run


# ==============================================================================================
# Target: list-routes
# Brief: Lists all routes in the API.
# Usage: Run the command 'make list-routes'.
# ==============================================================================================
.PHONY: list-routes
list-routes: .cmd-exists-go .clear-screen
ifneq ($(RUNNING_IN_DOCKER), true)
	@ENV_FILE=.env.dev $(MAKE) postgres
	@ENV_FILE=.env.dev $(MAKE) rabbitmq
	@ENV_FILE=.env.dev $(MAKE) redis
endif
	@go run ./cmd/routes/main.go


# ==============================================================================================
# Target: list-env
# Brief: Lists all environment variables in the environment file.
# Usage: Run the command 'make list-env ENV_FILE="<path>"'.
# Flags:
#  ENV_FILE: The path to the environment file.
# ==============================================================================================
.PHONY: list-env
list-env: .cmd-exists-go .clear-screen .check-env-file
	@go run ./cmd/env/*.go -e $(ENV_FILE)


# ==============================================================================================
# Target: mail
# Brief: Sends a test email.
# Usage: Run the command 'make mail ENV_FILE="<path>" TO="<email>" SUBJECT="<subject>" BODY="<body>" NAME="<name>"'.
# Flags:
# 	ENV_FILE: The path to the environment file.
#   TO: The email address to send the email to.
#   SUBJECT: The subject of the email.
#   NAME: The name of the recipient.
# ==============================================================================================
.PHONY: mail
mail: .cmd-exists-go .clear-screen .check-env-file
	@if [ -z "$(TO)" ]; then \
		echo "Error: expected TO"; \
		exit 1; \
	fi;

	@if [ -z "$(SUBJECT)" ]; then \
		echo "Error: expected SUBJECT"; \
		exit 1; \
	fi;

	@if [ -z "$(NAME)" ]; then \
		echo "Error: expected NAME"; \
		exit 1; \
	fi;

	@go run ./cmd/mail/*.go -e $(ENV_FILE) -t "$(TO)" -s "$(SUBJECT)" -n "$(NAME)" -p "$(PROVIDER)"


# ==============================================================================================
# Target: build
# Brief: Builds the API server.
# Usage: Run the command 'make build'.
# ==============================================================================================
.PHONY: build
build: .cmd-exists-go .clear-screen
	@echo "Generating build for $(APP_NAME)..."
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@echo "Build was generated at '$(BUILD_DIR)/$(APP_NAME)'."


# ==============================================================================================
# Target: lint
# Brief: Runs the linter (golangci-lint) in the project.
# Usage: Run the command 'make lint'.
# ==============================================================================================
.PHONY: lint
lint: .cmd-exists-docker .cmd-exists-sh .clear-screen
	@sh -c "./scripts/linter.sh"


# ==============================================================================================
# Target: clear
# Brief: Clears the auto-generated files and folders, such as the Docker volumes.
# Usage: Run the command 'make clear'.
# ==============================================================================================
.PHONY: clear
clear: .cmd-exists-sh .clear-screen
	@rm -rf $(BUILD_DIR)
	@rm -rf ./coverage.out
	@rm -rf ./coverage.html
	@rm -rf ./tmp


# ==============================================================================================
# Target: tidy
# Brief: Installs the dependencies and syncs the vendor folder.
# Usage: Run the command 'make tidy'.
# ==============================================================================================
.PHONE: tidy
tidy: .cmd-exists-go .clear-screen
	@go mod tidy
	@go mod vendor


# ==============================================================================================
# Target: test-unit
# Brief: Runs unit tests.
# Usage:
#   FLAG=<watch|cover> make test-unit # Runs tests normally without any special flags
#   COUNT=<count> make-test-unit # Runs tests <count> times
# Environment variables:
#   COUNT: Number of times to run each test. Cannot be used with other flags.
#   FLAG:  Defines the flag to run the tests with. Possible values: "cover", "watch".
# ==============================================================================================
.PHONY: test-unit
test-unit: .cmd-exists-go .clear-screen
	@TEST_MODE=unit COUNT=$(COUNT) FLAG=$(FLAG) \
	sh -c "./scripts/test.sh"


# ==============================================================================================
# Target: test-integration
# Brief: Runs integration tests.
# Usage: Run the command 'make test-unit [FLAGS="<flags>"]'.
# Usage:
#   FLAG=<watch|cover> make test-integration # Runs tests normally without any special flags
#   COUNT=<count> make-integration # Runs tests <count> times
## Environment variables:
#   COUNT: Number of times to run each test. Cannot be used with other flags.
#   FLAG:  Defines the flag to run the tests with. Possible values: "cover", "watch".
# ==============================================================================================
.PHONY: test-integration
test-integration: .cmd-exists-go .clear-screen .prepare-test
	@TEST_MODE=integration COUNT=$(COUNT) FLAG=$(FLAG) \
	sh -c "./scripts/test.sh"


# ==============================================================================================
# Target: test-e2e
# Brief: Runs end-to-end tests.
# Usage: Run the command 'make test-e2e [FLAGS="<flags>"]'.
# Usage:
#   FLAG=<watch|cover> make test-e2e # Runs tests normally without any special flags
#   COUNT=<count> make-e2e # Runs tests <count> times
# Environment variables:
#   COUNT: Number of times to run each test. Cannot be used with other flags.
#   FLAG:  Defines the flag to run the tests with. Possible values: "cover", "watch".
# ==============================================================================================
.PHONY: test-e2e
test-e2e: .cmd-exists-go .clear-screen .prepare-test
	@TEST_MODE=e2e COUNT=$(COUNT) FLAG=$(FLAG) \
	sh -c "./scripts/test.sh"


# ==============================================================================================
# Target: test
# Brief: Runs all tests.
# Usage:
#   FLAG=<watch|cover> make test # Runs tests normally without any special flags
#   COUNT=<count> make test # Runs tests <count> times
# Environment variables:
#   COUNT: Number of times to run each test. Cannot be used with other flags.
#   FLAG:  Defines the flag to run the tests with. Possible values: "cover", "watch".
# ==============================================================================================
.PHONY: test
test: .cmd-exists-go .clear-screen .prepare-test
	@TEST_MODE=all COUNT=$(COUNT) FLAG=$(FLAG) \
	sh -c "./scripts/test.sh"


# ==============================================================================================
# Target: test-ci
# Brief: Runs all tests in CI mode.
# Usage: Run the command 'make test-ci'.
# ==============================================================================================
.PHONY: test-ci
test-ci: .cmd-exists-go .clear-screen .prepare-test
	@TEST_MODE=all go test -v ./...


# ==============================================================================================
# Target: create-migration
# Brief: Creates a new migration file with the given name with the current timestamp.
# Usage: Run the command 'make create-migration'.
# ==============================================================================================
.PHONY: create-migration
create-migration: .cmd-exists-go .clear-screen
	@read -p "Enter the name of the migration file: " NAME; \
	sh -c "./scripts/create_migration.sh $$NAME"


# ==============================================================================================
# Target: migrate
# Brief: Runs the migrations.
# Usage: Run the command 'make migrate ENV_FILE="<env_file>" MIGRATION="<migration>" [VERSION="<version>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# 	MIGRATION: The migration to be run. Options: up, down, drop, force, steps, version, reset.
# 	VERSION: The version to be migrated.
# ==============================================================================================
.PHONY: migrate
migrate: .cmd-exists-docker .clear-screen .check-env-file
	@ENV_FILE=$(ENV_FILE) $(MAKE) postgres

	@if [ -z "$(MIGRATION)" ]; then \
		echo "Error: expected MIGRATION"; \
		exit 1; \
	fi;

	@case "$(MIGRATION)" in \
		up|down|force|drop|steps|version|reset) \
			if [ "$(MIGRATION)" = "steps" ] && [ -z "$(VERSION)" ]; then \
				echo "Error: expected VERSION"; \
				exit 1; \
			fi; \
			go run ./cmd/migrate/*.go -e "$(ENV_FILE)" "$(MIGRATION)" "$(VERSION)" ;; \
		*) \
			echo "Error: expected [up|down|force|drop|steps|version|reset]"; \
			exit 1; \
			;; \
	esac


# ==============================================================================================
# Target: postgres
# Brief: Runs the postgres in a docker container.
# Usage: Run the command 'make postgres [ENV_FILE="<env_file>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: postgres
postgres: .cmd-exists-docker .clear-screen .check-env-file
	@if [ "$(ENV_FILE)" = ".env.test" ]; then \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=psql_test; \
	else \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=psql; \
	fi;


# ==============================================================================================
# Target: rabbitmq
# Brief: Runs the rabbitmq in a docker container.
# Usage: Run the command 'make rabbitmq [ENV_FILE="<env_file>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: rabbitmq
rabbitmq: .cmd-exists-docker .clear-screen .check-env-file
	@if [ "$(ENV_FILE)" = ".env.test" ]; then \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=rabbitmq_test; \
	else \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=rabbitmq; \
	fi;


# ==============================================================================================
# Target: redis
# Brief: Runs the redis in a docker container.
# Usage: Run the command 'make redis [ENV_FILE="<env_file>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: redis
redis: .cmd-exists-docker .clear-screen .check-env-file
	@if [ "$(ENV_FILE)" = ".env.test" ]; then \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=redis_test; \
	else \
		$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=redis; \
	fi;


# ==============================================================================================
# Target: grafana
# Brief: Runs the grafana in a docker container.
# Usage: Run the command 'make grafana [ENV_FILE="<env_file>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: grafana
grafana: .cmd-exists-docker .clear-screen .check-env-file
	@$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=grafana


# ==============================================================================================
# Target: mock
# Brief: Generates mocks for all interfaces in the project, excluding the vendor folder.
# Usage: Run the command 'make mock'.
# ==============================================================================================
.PHONY: mock
mock: .cmd-exists-go
	@rm -rf ./testutils/mocks
	@sh -c "./scripts/mock.sh ./internal" &	\
	sh -c "./scripts/mock.sh ./pkg"


# ==============================================================================================
# Target: docker-dev
# Brief: Runs the API container in development mode.
# Usage: Run the command 'make docker-dev'.
# ==============================================================================================
.PHONY: docker-dev
docker-dev: .cmd-exists-docker .clear-screen
	@ENV_FILE=.env.dev \
	$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=api


# ==============================================================================================
# Target: docker-prod
# Brief: Runs the API container in production mode.
# Usage: Run the command 'make docker-prod'.
# ==============================================================================================
.PHONY: docker-prod
docker-prod: .cmd-exists-docker .clear-screen
	@ENV_FILE=.env.prod \
	$(MAKE) .docker COMMAND=up FLAG=-d SERVICE=api


# ==============================================================================================
# Target: docker-stop
# Brief: Stops the all containers in the docker-compose file.
# Usage: Run the command 'make docker-stop'.
# ==============================================================================================
.PHONY: docker-stop
docker-stop: .cmd-exists-docker .clear-screen
	@ENV_FILE=.env.dev \
	$(MAKE) .docker COMMAND=stop


# ==============================================================================================
# Target: docker-kill
# Brief: Kills the all containers in the docker-compose file.
# Usage: Run the command 'make docker-kill'.
# ==============================================================================================
.PHONY: docker-kill
docker-kill: .cmd-exists-docker .clear-screen
	@ENV_FILE=.env.dev \
	$(MAKE) .docker COMMAND=kill


# ==============================================================================================
# Target: docker-list-env
# Brief: Lists the environment variables in the given file.
# Usage: Run the command 'make docker-list-env ENV_FILE="<path>"'.
# Flags:
# 	ENV_FILE: The path to the environment file.
# ==============================================================================================
.PHONY: docker-list-env
docker-list-env: .cmd-exists-docker .clear-screen .check-env-file
	@WORKDIR=$(WORKDIR)  \
	AIRVERSION=$(AIRVERSION) \
	docker compose \
	--env-file "$(ENV_FILE)" \
	run \
	--name $(APP_NAME)-list-env \
	--rm \
	-e ENV_FILE=$(ENV_FILE) \
	api \
	make list-env


# ==============================================================================================
# Target: .docker
# Brief: This is a helper target to run docker-compose commands.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.PHONY: .docker
.docker: .cmd-exists-docker .clear-screen .check-env-file
	@RUNNING_IN_DOCKER=true \
	BUILD_DIR=$(BUILD_DIR) \
	WORKDIR=$(WORKDIR) \
	AIRVERSION=$(AIRVERSION) \
	docker compose --env-file $(ENV_FILE) $(COMMAND) $(FLAG) $(SERVICE)


# ==============================================================================================
# Target: .cmd-exists-%
# Brief: This is a helper target to check if a command exists. It will exit with code 1 if it does not.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)


# ==============================================================================================
# Target: .clear-screen
# Brief: This is a helper target to clear the terminal screen.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.clear-screen:
	@printf "\033c"


# ==============================================================================================
# Target: .check-env-file
# Brief: This is a helper target to check if the ENV_FILE is valid. It will exit with code 1 if it is not.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.check-env-file:
	@if [ -z "$(ENV_FILE)" ]; then \
		echo "Error: expected ENV_FILE"; \
		exit 1; \
	fi;
	
	@case "$(ENV_FILE)" in \
		.env.dev|.env.test|.env.prod) ;; \
		*) \
			echo "Error: expected .env.dev, .env.test or .env.prod"; \
			exit 1; \
			;; \
	esac


# ==============================================================================================
# Target: .prepare-test
# Brief: This is a helper target to prepare the test environment.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.prepare-test: .cmd-exists-go .cmd-exists-docker .clear-screen
	@ENV_FILE=.env.test $(MAKE) rabbitmq
	@ENV_FILE=.env.test $(MAKE) postgres
	@ENV_FILE=.env.test $(MAKE) redis
	@go run ./cmd/migrate/*.go -e .env.test reset


# ==============================================================================================
# DO NOT EDIT BELOW THIS LINE
# ==============================================================================================
MAKE = make --no-print-directory
BUILD_DIR = bin/build