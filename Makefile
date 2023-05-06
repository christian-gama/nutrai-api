PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory
APP_NAME = nutrai-api

# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/go/src/github.com/christian-gama/nutrai-api

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

# ==============================================================================================
# Target: lint
# Brief: This target is used to lint the project.
# Usage: Run the command 'make lint'.
# ==============================================================================================
.PHONY: lint
lint: .cmd-exists-docker .cmd-exists-sh .clear-screen
	@sh -c "./scripts/linter.sh"


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
#  --watch: Watch for changes and run tests.
# ==============================================================================================
.PHONY: test-integration
test-integration: .cmd-exists-go .clear-screen
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
test: .cmd-exists-go .clear-screen
	@TEST_MODE=all sh -c "./scripts/test.sh $(FLAGS)"


# ==============================================================================================
# Target: test-ci
# Brief: This target is used to run all tests on CI.
# Usage: Run the command 'make test-ci'.
# ==============================================================================================
.PHONY: test-ci
test-ci: .cmd-exists-go .clear-screen
	@TEST_MODE=all go test -v ./...


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

