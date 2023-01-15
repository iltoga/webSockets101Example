#!/bin/bash

## in alternative you can use `npm-run-all`
# npm install --save-dev npm-run-all
## Then you need to add the following scripts to your package.json file:
# "scripts": {
#     "start:client": "cd client && npm start",
#     "start:server": "cd server && go run main.go",
#     "start": "npm-run-all --parallel start:client start:server",
#     ...
# },
## When you run npm start in the root of your project, it will run the client and server in parallel.
## If you want to monitor the processes that are run with npm-run-all, you can use a tool like forever or pm2.
# npm install -g pm2
# pm2 start npm -- start
## This command will run the npm start command and will keep it running in the background.
## To log to file
# pm2 start npm -- start > output.log
## to see logs realtime
# pm2 logs

# Start the React client
cd client
npm start &

# Start the Golang server
cd ../server
go run main.go &

# Keep the script running
wait