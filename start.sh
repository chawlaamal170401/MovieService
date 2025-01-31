#!/bin/bash

# Start the server in the background
go run cmd/server/main.go &

# Start the client in the background
go run cmd/client/main.go &

# Wait for both processes to finish
wait
