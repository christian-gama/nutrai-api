#!/bin/bash
# ==============================================================================================
# Title:    linter.sh
# Brief:    Run golangci-lint in a Docker container to lint the source code.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-05
# Usage:    ./scripts/linter.sh
# ==============================================================================================

CACHE_DIR="$PWD/.cache/linter"
DEFAULT_CONFIG="$HOME/.golangci.yml"
mkdir -p "$CACHE_DIR"

docker run \
--rm -t \
--user "$(id -u):$(id -g)" \
-v "$CACHE_DIR:/.cache" \
-v "$PWD:/app" \
--workdir /app \
golangci/golangci-lint:v1.47.3 \
golangci-lint run --config .golangci.yml "$@"