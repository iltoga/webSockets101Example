package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var logToFile bool

func init() {
	flag.BoolVar(&logToFile, "log-to-file", false, "Log to file instead of stderr")
	flag.Parse()
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	if logToFile {
		file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		logger.Out = file
	}

	logger.Debug("Client Connected")
	err = ws.WriteMessage(1, []byte("Hello client you've connected!"))
	if err != nil {
		logger.Error(err)
	}
	reader(ws, logger)
}

func reader(conn *websocket.Conn, logger *logrus.Logger) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}

		// log the incoming message
		logger.Debug("incoming message: " + string(p))

		// respond to the client with the message in uppercase
		response := []byte(strings.ToUpper(string(p)))
		if err := conn.WriteMessage(messageType, response); err != nil {
			logger.Error(err)
			return
		}

		// log the message sent
		logger.Info("sent message: " + string(response))
	}
}
