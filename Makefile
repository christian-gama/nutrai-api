PWD = $(shell pwd)
GENERATED_DIR ?= $(PWD)/.generated
CACHE_DIR ?= $(PWD)/.cache
BUILD_DIR ?= $(GENERATED_DIR)/build
MAKE = make --no-print-directory
APP_NAME = nutrai-api

# WORKDIR is used to set the working directory for Dockerfile builds.
export WORKDIR=/go/src/github.com/christian-gama/nutrai-api

-include $(ENV_FILE)

.PHONY: init
init: cmd-exists-git cmd-exists-go cmd-exists-docker cmd-exists-sh clear-screen
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Git Hooks configured."

	@chmod +x ./scripts/*.sh
	@echo "Scripts configured."


.PHONY: lint
lint: cmd-exists-docker cmd-exists-sh clear-screen
	@sh -c "./scripts/linter.sh"


.PHONY: build
build: cmd-exists-go clear-screen
	@echo "Generating build for $(APP_NAME)..."
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/*.go
	@echo "Build was generated at '$(BUILD_DIR)/$(APP_NAME)'."


.PHONE: tidy
tidy: cmd-exists-go clear-screen
	@go mod tidy
	@go mod vendor


.PHONY: clear-screen
clear-screen:
	@printf "\033c"


cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)