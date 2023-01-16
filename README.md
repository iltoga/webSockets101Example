
# webSockets101Example

Example of application using websockets for client-server real-time communitcation. Server side in golang and client in React js

There are two applications in the relative directories:
- client: React functional application (in typescript) that implements the frontend
- server: Golang application that implements websocket server

## Client
In the component, it uses the "useEffect" hook to initiate a WebSocket connection to a server running on 'ws://localhost:8080/ws' upon the component mounting. It sends a message "helloheee!" when the connection is opened and logs any incoming messages to the console.

## Server
This is a Go program that creates a WebSocket server that listens on port 8080. It uses the "net/http" and "github.com/gorilla/websocket" libraries.

In the main function, it sets up an http handler for the path "/ws" that calls the "wsEndpoint" function. It starts a server on port 8080 using the "http.ListenAndServe" function.

  

The "wsEndpoint" function accepts an http response writer and request as its parameters. It uses the "websocket.Upgrader" struct to upgrade the connection to a WebSocket connection, and the connection is set to allow any origin. It logs that a client has connected, sends a message "Hello client you've connected!" to the client, and starts a reader function to read in messages and send them back to the client.

The "reader" function continuously listens for incoming messages on the WebSocket connection, logs the incoming message, and sends it back to the client.

## How to run
```
chmod +x run.sh
./run.sh
```
In alternative you can use `npm-run-all`
```
npm install --save-dev npm-run-all
```
Then you need to add the following scripts to your package.json file:
```
"scripts": {
"start:client": "cd client && npm start",
"start:server": "cd server && go run main.go",
"start": "npm-run-all --parallel start:client start:server",
...
},
```
When you run npm start in the root of your project, it will run the client and server in parallel.

If you want to monitor the processes that are run with npm-run-all, you can use a tool like forever or pm2.
```
npm install -g pm2
pm2 start npm -- start
```
This command will run the npm start command and will keep it running in the background.

To log to file
```
pm2 start npm -- start > output.log
```
To see logs realtime
```
pm2 logs
```