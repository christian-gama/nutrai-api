#!/bin/bash
# ==============================================================================================
# Title:    docker_log.sh
# Brief:    Print the logs from the Docker container of the selected service. If no service is
#           selected, a menu will be displayed.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-06-06
# Usage:    ./scripts/docker_log.sh [service]
# ==============================================================================================

declare -A options
options=(
    [1]="grafana"
    [2]="prometheus"
    [3]="api"
    [4]="psql"
    [5]="redis"
    [6]="rabbitmq"
    [9]="exit"
)

declare -A service_map
service_map=(
    ["grafana"]=1
    ["prometheus"]=2
    ["api"]=3
    ["psql"]=4
    ["redis"]=5
    ["rabbitmq"]=6
)

function menu {
    for option in $(echo ${!options[@]} | tr " " "\n" | sort -n); do
        echo "$option) ${options[$option]}"
    done
}


function docker_compose {
    docker-compose --env-file .env.dev logs -f --tail=100 "${options[$1]}"
}

function log {
    if [[ "$1" =~ ^[0-9]+$ ]]; then
        case "$1" in
            9) exit 0 ;;
            1|2|3|4|5|6) docker_compose "$1" ;;
            *) echo "Invalid option" ;;
        esac
    else
        docker_compose "${service_map[$1]}"
    fi
}

function clear_screen {
    clear
}

if [ $# -eq 0 ]; then
    while true; do
        clear_screen
        menu
        read -p "Select an option: " option
        log "$option"
    done
else
    log "$1"
fi
