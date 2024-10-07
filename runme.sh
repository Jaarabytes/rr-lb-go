#!/bin/sh

URL="http://localhost:8080"
NUM_REQUESTS=20
LOG_FILE="runme.log"

# Simple bash script that sends 20 POST requests to the HTTP server
#  The load balancer tries to hash the random string and write the hash to the database
#  It is hashed such that data is easily found.


echo "Sending requests SOON"
for i in $(seq 1 $NUM_REQUESTS)
do
    RANDOM_STRING=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
    curl -X POST -d "$RANDOM_STRING" $URL
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $RANDOM_STRING" >> "$LOG_FILE"
    # The server isn't that strong. Give it some room
    # sleep 0.5
done

echo "All requests sent"
