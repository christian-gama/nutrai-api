#!/bin/bash
# ==============================================================================================
# Title:    create_env.sh
# Brief:    Create .env.dev, .env.test, and .env.prod files if they don't exist.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-07
# Usage:    ./scripts/create_test.sh
# ==============================================================================================

create_env_file() {
    env_name="$1"
    env_file=".env.$env_name"

    if [ ! -f "$env_file" ]; then
        cp .env.example "$env_file"
        sed -i "s/APP_ENV=.*/APP_ENV=$env_name/" "$env_file"
        
        if [ "$env_name" == "dev" ]; then
            sed -i "s/DB_PORT=.*/DB_PORT=5433/" "$env_file"
        fi

        if [ "$env_name" == "test" ]; then
            sed -i "s/DB_PORT=.*/DB_PORT=5434/" "$env_file"
        fi

        echo "Created $env_file"
    else
        echo "$env_file already exists. Skipping."
    fi
}

# Create .env.dev, .env.test, and .env.prod files
create_env_file "dev"
create_env_file "test"
create_env_file "prod"
