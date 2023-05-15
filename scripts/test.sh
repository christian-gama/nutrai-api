#!/bin/bash
# ==============================================================================================
# Title:  test.sh
# Brief:  Run tests
# Author: christiangama.dev@gmail.com
# Creation: 2023-05-06
# Usage:  ./scripts/test.sh [FLAGS=<flags>]
# Flags:
#   --watch: Run tests in watch mode
#   --cover: Run tests with coverage
# ==============================================================================================

FLAG=$1

if [ "$FLAG" = "--cover" ]; then
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    -- \
    -coverprofile=coverage.out \
    ./... && \
    go tool cover -html=coverage.out -o coverage.html
else
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    $FLAG \
    -- \
    ./...
fi