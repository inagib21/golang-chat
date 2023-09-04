package websocket

import "fmt"

// * Pool represents a WebSocket client pool.
type Pool struct {
	Register   chan *Client     //* Channel for registering clients
	Unregister chan *Client     //* Channel for unregistering clients
	Clients    map[*Client]bool //* Map to store connected clients
	Broadcast  chan Message     //* Channel for broadcasting messages to clients
}

// * NewPool creates a new WebSocket client pool.
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// * Start begins the WebSocket client pool and handles client registration, unregistration, and broadcasting messages.
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			// Register a new client and notify all clients about the new user.
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			// Unregister a client and notify all clients about the user's disconnection.
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			// Broadcast a message to all connected clients in the pool.
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
