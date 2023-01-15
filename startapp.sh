#!/bin/bash

# Start the React client
cd client
npm start &

# Start the Golang server
cd ../server
go run main.go &

# Keep the script running
wait

