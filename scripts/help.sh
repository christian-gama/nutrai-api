#!/bin/bash
# ==============================================================================================
# Title:    help.sh
# Brief:    List all Makefile targets and their brief description.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-25
# Usage:    ./scripts/help.sh
# ==============================================================================================

echo -e "\e[1;33m=============================================\e[0m"
echo -e "\e[1;33mListing all Makefile targets (\e[1;32mtarget\e[0m - \e[1;34mbrief\e[0m)\e[0m"
echo -e "\e[1;33mTo run a target, use: make \e[1;32mtarget\e[0m\e[0m"
echo -e "\e[1;33m=============================================\e[0m"

grep -E '^# Target:|^# Brief:' Makefile | while read -r line
do
    if [[ $line == "# Target:"* ]]; then
        target=${line#*:}
        target_trimmed=$(echo $target | xargs)  # Trimming leading/trailing whitespaces

        # Ignore targets that start with a dot
        if [[ $target_trimmed != .* ]]; then
            printf "\e[1;32m%s\e[0m" "$target"
            read -r line
            if [[ $line == "# Brief:"* ]]; then
                printf " - \e[1;34m%s\e[0m\n" "${line#*:}"
            fi
        fi
    fi
done

echo -e "\n"
