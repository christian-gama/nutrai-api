#!/bin/bash
# ==============================================================================================
# Title:    manage_env.sh
# Brief:    Create .env.dev, .env.test, and .env.prod files if they don't exist. If they exist,
#           update them with the latest keys from .env.example. If a key is not present in 
#           .env.example, remove it from the environment specific file.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-07
# Usage:    ./scripts/manage_env.sh
# ==============================================================================================

manage_env_file() {
    env_name="$1"
    env_file=".env.$env_name"
    example_file=".env.example"

    if [ -f "$env_file" ]; then
        temp_file=$(mktemp)
        copy_and_update_env_file "$example_file" "$env_file" "$temp_file"
        remove_unused_keys "$example_file" "$temp_file"
        move_temp_file "$temp_file" "$env_file"
    else
        cp "$example_file" "$env_file"
    fi

    update_environment_specific_variables "$env_name" "$env_file"
    display_result "$env_file"
}

copy_and_update_env_file() {
    example_file="$1"
    env_file="$2"
    temp_file="$3"

    while IFS= read -r line
    do
        if is_comment_line "$line"; then
            continue
        fi

        key=$(extract_key "$line")
        if grep -q "^$key=" "$env_file"; then
            grep "^$key=" "$env_file" >> "$temp_file"
        else
            echo "$line" >> "$temp_file"
        fi
    done < "$example_file"
}

remove_unused_keys() {
    example_file="$1"
    temp_file="$2"

    while IFS= read -r line
    do
        if is_comment_line "$line"; then
            continue
        fi

        key=$(extract_key "$line")
        if ! grep -q "^$key=" "$example_file"; then
            sed -i "/^$key=/d" "$temp_file"
        fi
    done < "$env_file"
}

update_environment_specific_variables() {
    env_name="$1"
    env_file="$2"

    sed -i "s/APP_ENV=.*/APP_ENV=\"$env_name\"/" "$env_file"

    case "$env_name" in
        "dev")
            sed -i "s/DB_PORT=.*/DB_PORT=5433/" "$env_file"
            sed -i "s/RABBITMQ_PORT=.*/RABBITMQ_PORT=5673/" "$env_file"
            sed -i "s/CONFIG_LOG_LEVEL=.*/CONFIG_LOG_LEVEL=\"debug\"/" "$env_file"
            ;;
        "test")
            sed -i "s/DB_PORT=.*/DB_PORT=5434/" "$env_file"
            sed -i "s/RABBITMQ_PORT=.*/RABBITMQ_PORT=5674/" "$env_file"
            sed -i "s/CONFIG_LOG_LEVEL=.*/CONFIG_LOG_LEVEL=\"panic\"/" "$env_file"
            ;;
        "prod")
            sed -i "s/CONFIG_GLOBAL_RATE_LIMIT=.*/CONFIG_GLOBAL_RATE_LIMIT=180/" "$env_file"
            sed -i "s/CONFIG_DEBUG=.*/CONFIG_DEBUG=false/" "$env_file"
            ;;
    esac
}

move_temp_file() {
    temp_file="$1"
    env_file="$2"

    mv "$temp_file" "$env_file"
}

is_comment_line() {
    line="$1"
    [[ "$line" =~ ^#.*$ ]]
}

extract_key() {
    line="$1"
    echo "$line" | cut -d '=' -f 1
}

display_result() {
    env_file="$1"
    echo "Done with $env_file"
}

# Create .env.dev, .env.test, and .env.prod files
manage_env_file "dev"
manage_env_file "test"
manage_env_file "prod"
