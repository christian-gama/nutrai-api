#!/bin/bash
# ==============================================================================================
# Title:    help.sh
# Brief:    List all Makefile targets and their brief description.
# Author:   christiangama.dev@gmail.com
# Creation: 2023-05-25
# Usage:    ./scripts/help.sh
# ==============================================================================================

echo -e "===================================================="
echo -e "Listing all Makefile targets (\e[1;32mtarget\e[0m - \e[1;34mbrief\e[0m)\e[0m"
echo -e "To run a target, run the command: make \e[1;32mtarget\e[0m"
echo -e "===================================================="

grep -E '^# Target:|^# Brief:' Makefile | while read -r line
do
    if [[ $line == "# Target:"* ]]; then
        target=${line#*:}
        target_trimmed=$(echo $target | xargs)  # Trimming leading/trailing whitespaces

        # Ignore targets that start with a dot
        if [[ $target_trimmed != .* ]]; then
            printf "\e[1;32m%s\e[0m" "$target_trimmed"
            read -r line
            if [[ $line == "# Brief:"* ]]; then
                brief=${line#*:}
                brief_trimmed=$(echo $brief | xargs)  # Trimming leading/trailing whitespaces
                printf " - \e[1;34m%s\e[0m\n" "$brief_trimmed"
            fi
        fi
    fi
done
