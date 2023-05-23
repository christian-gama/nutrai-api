#!/bin/bash
# ==============================================================================================
# Title:  test.sh
# Brief:  Run tests
# Author: christiangama.dev@gmail.com
# Creation: 2023-05-06
# Usage:
#   FLAG=<watch|cover> ./scripts/test.sh # Runs tests normally without any special flags
#   COUNT=<count> ./scripts/test.sh # Runs tests <count> times
#
# Environment variables:
#   COUNT: Number of times to run each test. Cannot be used with other flags.
#   FLAG:  Defines the flag to run the tests with. Possible values: "cover", "watch".
# ==============================================================================================

set -e

run_test() {
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    --$1 \
    -- \
    ./...
}

run_test_with_count() {
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    -- \
    -count=$1 \
    ./...
}

run_test_with_cover() {
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    -- \
    -coverprofile=coverage.out \
    ./... && \
    go tool cover -html=coverage.out -o coverage.html
}

run_test_plain() {
    go run gotest.tools/gotestsum@v1.10.0 \
    --format pkgname \
    --format-hide-empty-pkg \
    --hide-summary skipped \
    -- \
    ./...
}

if [[ -n "$COUNT" ]] && [[ -n "$FLAG" ]]; then
    echo "Error: COUNT and FLAG cannot be used together."
    exit 1
elif [[ -n "$COUNT" ]]; then
    if [[ $COUNT -gt 0 ]]; then
        run_test_with_count $COUNT
    else
        echo "Error: COUNT should be a positive integer."
        exit 1
    fi
elif [[ "$FLAG" == "cover" ]] || [[ "$FLAG" == "watch" ]]; then
    if [[ "$FLAG" == "cover" ]]; then
        run_test_with_cover
    elif [[ "$FLAG" == "watch" ]]; then
        run_test $FLAG
    fi
else
    run_test_plain
fi