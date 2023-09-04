package main

import (
	"fmt"
	"net/http"

	"github.com/inagib21/golang-chat/pkg/websocket"
)

// serveWS handles WebSocket requests and initializes a client connection.
func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket endpoint reached")

	// Upgrade the HTTP connection to WebSocket.
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// Create a new WebSocket client and register it with the pool.
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client

	// Start reading messages from the WebSocket client.
	client.Read()
}

// setupRoutes initializes WebSocket pool and sets up WebSocket route.
func setupRoutes() {
	// Create a new WebSocket pool.
	pool := websocket.NewPool()

	// Start the WebSocket pool in a separate goroutine.
	go pool.Start()

	// Define a WebSocket route "/ws" and handle WebSocket requests using serveWS function.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	// Print a message indicating the start of the application.
	fmt.Println("Nagib's full stack chat project")

	// Set up WebSocket routes and start the HTTP server on port 9000.
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
