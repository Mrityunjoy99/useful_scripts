#!/bin/bash

# Define an array of ports to check
ports=("$@")

# Loop through the ports and attempt to find the processes running on them
for port in "${ports[@]}"; do
  # Use lsof to find the process using the port
  pid=$(lsof -t -i :$port)

  if [ -n "$pid" ]; then
    echo "Killing process on port $port (PID: $pid)"
    kill $pid
  else
    echo "No process found on port $port"
  fi
done