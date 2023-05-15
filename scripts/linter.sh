#!/bin/bash
# ==============================================================================================
# Title:    linter.sh
# Brief:    Run golangci-lint in a Docker container to lint the source code.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-05
# Usage:    ./scripts/linter.sh
# ==============================================================================================

DEFAULT_CONFIG=".golangci.yml"

echo "Running golangci-lint..."
go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2 run --config "$DEFAULT_CONFIG" ./...