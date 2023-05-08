#!/bin/bash
# ==============================================================================================
# Title:    wait_for_db.sh
# Brief:    Wait for Database to be up and running.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-07
# Usage:    ./scripts/wait_for_db.sh <container_name>
# ==============================================================================================

container_name="$1"

is_first_time=true

until docker inspect --format "{{.State.Health.Status}}" $container_name | grep "healthy" > /dev/null; do
  if $is_first_time; then
    >&2 echo "Connecting to $container_name - the first time may take a while"
    is_first_time=false
  else
    >&2 echo "Waiting for $container_name to be up and running..."
  fi
  sleep 2
done