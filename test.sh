#!/bin/bash

# URL of the web application
url="http://94.237.59.185:43871/question1/"

# Maximum number of requests before rate limiting is expected
max_requests=100

# Time to wait between requests (in seconds)
wait_time=1

# Make requests until rate limiting is hit
for ((i=1; i<=$max_requests; i++)); do
  response=$(curl -s -o /dev/null -w "%{http_code}" $url)

  # Check if the response indicates rate limiting (HTTP 429)
  if [[ $response -eq 429 ]]; then
    echo "Rate limit hit after $i requests. Waiting for $wait_time seconds..."
    sleep $wait_time
  fi
done
