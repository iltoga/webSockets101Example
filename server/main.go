package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.ListenAndServe(":8000", nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	for {
		// Read in a new message as a JSON object and map it to a Message object
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		// Send the newly received message to the broadcast channel
		fmt.Printf("Received message: %s\n", msg)
	}
}
