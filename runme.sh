#!/bin/sh

URL="http://localhost:8080"
NUM_REQUESTS=20

# Simple bash script that sends 20 POST requests to the HTTP server
#  The load balancer tries to hash the random string and write the hash to the database
#  It is hashed such that data is easily found.


echo "Sending requests SOON"
for i in $(seq 1 $NUM_REQUESTS)
do
    RANDOM_STRING=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
    echo "Random string is $RANDOM_STRING"    
    curl -X POST -d "$RANDOM_STRING" $URL
    echo ""
    echo "Request $i sent"
    echo "" 
    # The server isn't that strong. Give it some room
    sleep 0.5
done

echo "All requests sent"
