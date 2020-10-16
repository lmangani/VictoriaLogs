#!/bin/bash
## Test Metrics for Loki/cLoki/VictoriaLogs

        DOMAINS=("up" "down" "left" "right")

        for i in `seq 1 10`;
        do
                TIME=$(date --utc +%FT%T.%3NZ)
                RANDOM=$$$(date +%s)
                NAME=${DOMAINS[$RANDOM % ${#DOMAINS[@]}]}
                echo "$NAME, $TIME"
                curl  --header "Content-Type: application/json"  --request POST \
                        --data '{"streams": [{"labels": "{level=\"WARN\",name=\"'"$NAME"'\"}","entries": [{"ts": "'"$TIME"'", "line": "level=info string='"$RANDOM"'" }]}]}' \
                        'http://127.0.0.1:8480/insert/0/loki/api/v1/push' &
        done
