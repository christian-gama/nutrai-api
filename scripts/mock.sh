#!/bin/bash
# ==============================================================================================
# Title:    mock.sh
# Brief:    Mock file for tests.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-07
# Usage:    ./scripts/mock.sh <dir>
# ==============================================================================================

dir=$1

echo "Running mockery for $dir..."
go run github.com/vektra/mockery/v2@v2.20.0 \
	--all \
	--keeptree \
	--case underscore \
	--exported \
	--dir $dir \
	--quiet \
	--output ./testutils/mocks

