#!/bin/bash
# ==============================================================================================
# Title:    metric_key.sh
# Brief:    Create a file with the metric keys.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-06
# Usage:    ./scripts/metric_key.sh <env_file>
# ==============================================================================================

create_metric_key() {
  # Check if the file was provided
  if [[ -z "$1" ]]; then
      echo "Usage: $0 <env_file>"
      exit 1
  fi

  # Check if the file exists
  if [[ ! -f "$1" ]]; then
      echo "$1 does not exist."
      exit 1
  fi

  # Read the .env file and get the value of APP_API_KEY
  export $(grep -v '^#' $1 | xargs)

  # Create the metrics directory if it does not exist
  mkdir -p metrics

  # Write the value to the metrics/prometheus.key file
  echo "$APP_API_KEY" > metrics/prometheus/prometheus-$APP_ENV.key
}

create_metric_key ".env.dev"
create_metric_key ".env.prod"