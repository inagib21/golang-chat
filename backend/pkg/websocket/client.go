package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// * Client represents a WebSocket client.
type Client struct {
	ID   string          //* Unique identifier for the client
	Conn *websocket.Conn //* WebSocket connection for the client
	Pool *Pool           //* Reference to the WebSocket pool
	mu   sync.Mutex      //* Mutex for safe concurrent access to client properties
}

// * Message represents a WebSocket message.
type Message struct {
	Type int    `json:"type"` //* Message type (e.g., 1 for text message)
	Body string `json:"body"` //* Message body (text content)
}

// * Read continuously listens for incoming messages from the WebSocket client.
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//* Create a Message struct from the received data.
		message := Message{Type: messageType, Body: string(p)}

		//* Broadcast the received message to all clients in the pool.
		c.Pool.Broadcast <- message

		//* Print the received message for debugging purposes.
		fmt.Printf("Message Received: %+v\n", message)
	}
}
