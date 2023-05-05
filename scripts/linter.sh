#!/bin/bash
#
# Script to lint the code using golangci-lint docker image.
#
# Author: Christian Gama e Silva
# Email: christiangama.dev@gmail.com
# Date Created: 2023/05/04

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