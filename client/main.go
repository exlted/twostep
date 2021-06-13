package main

// Initial Code from : https://frontside.medium.com/websockets-in-50-lines-of-go-790214b9ed92

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func register() *websocket.Conn {
	// Establish a websocket connection with the server
	// The returned connection object can then be used to read and write messages from/to the server.
	connection, _, err := websocket.DefaultDialer.Dial("ws://server:8081/ws", nil)
	if connection == nil {
		log.Fatal(err)
	}
	return connection
}

func consume(connection *websocket.Conn) {
	// An infinite for loop reading from the websocket and printing the reveived message
	// Must be started in a go routine
	for {
		// I exit the loop upon error here so the program doesn't panic when one side closes the connection
		_, message, err := connection.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("server says >>", string(message))
	}
}

func send(connection *websocket.Conn, message string) {
	fmt.Println("sending:: ", message)
	connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func sendSomeMessages(connection *websocket.Conn) {
	send(connection, "Hello server, this is client")
	time.Sleep(2 * time.Second)
	send(connection, "What you up to?")
	time.Sleep(3 * time.Second)
	send(connection, "OK, bye now")
	time.Sleep(10 * time.Second)
}

func main() {
	// Call the /ws http endpoint on the server
	connection := register()
	// Start a go routine reading from the websocket
	go consume(connection)
	// Send some arbitrary messages to the socket for demo purposes
	sendSomeMessages(connection)
}
