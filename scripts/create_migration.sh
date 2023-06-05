#!/bin/bash
# ==============================================================================================
# Title:    create_migration.sh
# Brief:    Create a migration file with a timestamp and a name.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-06
# Usage:    ./scripts/create_migration.sh <name>
# ==============================================================================================

if [[ -z "$1" ]]; then
    echo "Usage: $0 <migration_file_name>"
    exit 1
fi

timestamp=$(date +%s)
migration_dir="migration"

up="$timestamp"_"$1.up".sql
down="$timestamp"_"$1.down".sql

touch "$migration_dir"/"$up"
touch "$migration_dir"/"$down"

file_content="BEGIN;
COMMIT;"

echo "$file_content" > "$migration_dir"/"$up"
echo "$file_content" > "$migration_dir"/"$down"
