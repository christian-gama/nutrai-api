PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory
APP_NAME = nutrai-api
AIRVERSION = v1.43.0

# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/usr/src/app

# ENV_FILE is used to load the environment variables from the .env file.
-include $(ENV_FILE)

# ==============================================================================================
# Target: init
# Brief: This target is used to initialize the project.
# Usage: Run the command 'make'.
# ==============================================================================================
.PHONY: init
init: .cmd-exists-git .cmd-exists-go .cmd-exists-docker .cmd-exists-sh .clear-screen
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Git Hooks configured."

	@chmod +x ./scripts/*.sh
	@echo "Scripts configured."
	@./scripts/create_env.sh


# ==============================================================================================
# Target: run
# Brief: This target is used to run the project.
# Usage: Run the command 'make run [ENV_FILE="<path>"]'.
# ==============================================================================================
.PHONY: run
run: .cmd-exists-go .clear-screen .check-env-file
ifneq ($(RUNNING_IN_DOCKER), true)
	@$(MAKE) postgres
	@sh ./scripts/wait_for_db.sh nutrai-psql
endif

ifeq ($(ENV_FILE), .env.prod)
	@$(MAKE) build
	@$(BUILD_DIR)/$(APP_NAME) -e $(ENV_FILE)
else
	@go run github.com/cosmtrek/air@$(AIRVERSION) -- -e $(ENV_FILE)
endif


# ==============================================================================================
# Target: build
# Brief: This target is used to build the project. It will generate a binary file at the BUILD_DIR.
# Usage: Run the command 'make build'.
# ==============================================================================================
.PHONY: build
build: .cmd-exists-go .clear-screen
	@echo "Generating build for $(APP_NAME)..."
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@echo "Build was generated at '$(BUILD_DIR)/$(APP_NAME)'."


# ==============================================================================================
# Target: lint
# Brief: This target is used to lint the project.
# Usage: Run the command 'make lint'.
# ==============================================================================================
.PHONY: lint
lint: .cmd-exists-docker .cmd-exists-sh .clear-screen
	@sh -c "./scripts/linter.sh"


# ==============================================================================================
# Target: clear
# Brief: This target is used to clear the project from any temporary files.
# Usage: Run the command 'make clear'.
# ==============================================================================================
.PHONY: clear
clear: .cmd-exists-sh .clear-screen
	@rm -rf $(GENERATED_DIR)
	@rm -rf $(CACHE_DIR)
	@rm -rf ./coverage.out
	@rm -rf ./coverage.html
	@rm -rf ./testutils/mocks
	@rm -rf ./tmp


# ==============================================================================================
# Target: tidy
# Brief: This target is used install dependencies and generate the vendor folder.
# Usage: Run the command 'make tidy'.
# ==============================================================================================
.PHONE: tidy
tidy: .cmd-exists-go .clear-screen
	@go mod tidy
	@go mod vendor


# ==============================================================================================
# Target: test-unit
# Brief: This target is used to run unit tests.
# Usage: Run the command 'make test-unit [FLAGS="<flags>"]'.
# Flags: 
#  --watch: Watch for changes and run tests.
# ==============================================================================================
.PHONY: test-unit
test-unit: .cmd-exists-go .clear-screen
	@TEST_MODE=unit sh -c "./scripts/test.sh $(FLAGS)"


# ==============================================================================================
# Target: test-integration
# Brief: This target is used to run integration tests.
# Usage: Run the command 'make test-unit [FLAGS="<flags>"]'.
# Flags: 
# 	--watch: Watch for changes and run tests.
# ==============================================================================================
.PHONY: test-integration
test-integration: .cmd-exists-go .clear-screen .prepare-test-db
	@TEST_MODE=integration sh -c "./scripts/test.sh $(FLAGS)"


# ==============================================================================================
# Target: test
# Brief: This target is used to run all tests.
# Usage: Run the command 'make test [FLAGS="<flags>"]'.
# Flags:
#  --watch: Watch for changes and run tests.
#  --cover: Run tests with coverage.
# ==============================================================================================
.PHONY: test
test: .cmd-exists-go .clear-screen .prepare-test-db
	@TEST_MODE=all sh -c "./scripts/test.sh $(FLAGS)"


# ==============================================================================================
# Target: test-ci
# Brief: This target is used to run all tests on CI.
# Usage: Run the command 'make test-ci'.
# ==============================================================================================
.PHONY: test-ci
test-ci: .cmd-exists-go .clear-screen .prepare-test-db
	@TEST_MODE=all go test -v ./...


# ==============================================================================================
# Target: create-migration
# Brief: This target is used to create a new migration file.
# Usage: Run the command 'make create-migration NAME="<name>"'.
# Args:
# 	NAME: The name of the migration.
# ==============================================================================================
.PHONY: create-migration
create-migration: .cmd-exists-go .clear-screen
	@sh -c "./scripts/create_migration.sh $(NAME)"


# ==============================================================================================
# Target: migrate
# Brief: This target is used to run migrations.
# Usage: Run the command 'make migrate ENV_FILE="<env_file>" MIGRATION="<migration>" [VERSION="<version>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# 	MIGRATION: The migration to be run. Options: up, down, drop, force, steps, version, reset.
# 	VERSION: The version to be migrated.
# ==============================================================================================
.PHONY: migrate
migrate: .cmd-exists-docker .clear-screen .check-env-file .prepare-db
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
# Brief: This target is used to run the postgres container.
# Usage: Run the command 'make postgres [ENV_FILE="<env_file>"]'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: postgres
postgres: .cmd-exists-docker .clear-screen .check-env-file
	@if [ "$(ENV_FILE)" = ".env.test" ]; then \
		WORKDIR=$(WORKDIR) AIRVERSION=$(AIRVERSION) docker compose --env-file $(ENV_FILE) up -d psql_test; \
	else \
		WORKDIR=$(WORKDIR) AIRVERSION=$(AIRVERSION) docker compose --env-file $(ENV_FILE) up -d psql; \
	fi;


# ==============================================================================================
# Target: mock
# Brief: This target is used to generate mocks.
# Usage: Run the command 'make mock'.
# ==============================================================================================
.PHONY: mock
mock: .cmd-exists-go
	@sh -c "./scripts/mock.sh ./internal"
	@sh -c "./scripts/mock.sh ./pkg"


# ==============================================================================================
# Target: docker-run
# Brief: This target is used to run the API container.
# Usage: Run the command 'make docker-run ENV_FILE="<env_file>"'.
# Args:
# 	ENV_FILE: The env file to be loaded.
# ==============================================================================================
.PHONY: docker-run
docker-run: .cmd-exists-docker .clear-screen .check-env-file
	@RUNNING_IN_DOCKER=true WORKDIR=$(WORKDIR) AIRVERSION=$(AIRVERSION) docker compose --env-file $(ENV_FILE) up -d api --build --force-recreate --remove-orphans


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

	@if [ "$(ENV_FILE)" != ".env.dev" ] && [ "$(ENV_FILE)" != ".env.prod" ] && [ "$(ENV_FILE)" != ".env.test" ]; then \
		echo "Error: expected .env.dev, .env.test or .env.prod"; \
		exit 1; \
	fi;


# ==============================================================================================
# Target: .prepare-test-db
# Brief: This is a helper target to prepare the test environment.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.prepare-test-db: .cmd-exists-go .cmd-exists-docker .clear-screen
	@WORKDIR=$(WORKDIR) AIRVERSION=$(AIRVERSION) docker compose --env-file .env.test up -d psql_test
	@sh ./scripts/wait_for_db.sh nutrai-psql-test
	@go run ./cmd/migrate/*.go -e .env.test reset


# ==============================================================================================
# Target: .prepare-db
# Brief: This is a helper target to prepare the database.
# Usage: It is not meant to be called directly, but by other targets.
# ==============================================================================================
.prepare-db: .cmd-exists-docker .clear-screen
	@WORKDIR=$(WORKDIR) AIRVERSION=$(AIRVERSION) docker compose --env-file $(ENV_FILE) up -d psql
	@sh ./scripts/wait_for_db.sh nutrai-psql
