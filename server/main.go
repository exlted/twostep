package main

// Initial Code from : https://frontside.medium.com/websockets-in-50-lines-of-go-790214b9ed92

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// The upgrader type that allows us to upgrade an HTTP connection
// to a websocket
var upgrader = websocket.Upgrader{}

func consume(connection *websocket.Conn) {
	// Same as on the client
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("client says >> ", string(message))
	}
}

func send(connection *websocket.Conn, message string) {
	fmt.Println("sending:: ", message)
	// Write a message to the given websocket.
	// websocket.TextMessage indicates that what we pass in,
	// is to be interpreted by the reveiver as a plain text message
	connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func register(w http.ResponseWriter, r *http.Request) {
	// This is the handler that is being called
	// when a client sends a standard HTTP request to the /ws endpoint

	// Take the HTTP connection and upgrade it to a websocket
	connection, _ := upgrader.Upgrade(w, r, nil)

	// For each client that connects we now start a go routine
	// reading message from the socket
	go consume(connection)

	// Send some random messages into the websocket
	// (again this happens for each client that connects)
	sendSomeMessages(connection)
}

func sendSomeMessages(connection *websocket.Conn) {
	time.Sleep(5 * time.Second)
	send(connection, "Hello from the server")
	time.Sleep(1 * time.Second)
	send(connection, "It's lovely out today, isn't it")
	time.Sleep(2 * time.Second)
	send(connection, "OK, See you.")
}

func main() {
	// Create an HTTP server and listen for connections to the endpoint /ws
	http.HandleFunc("/ws", register)

	http.ListenAndServe(":8081", nil)
}
