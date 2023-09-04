package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// * upgrader is a WebSocket connection upgrader with buffer size settings.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, //* Size of the read buffer
	WriteBufferSize: 1024, //* Size of the write buffer
}

// * Upgrade upgrades an HTTP connection to a WebSocket connection.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// CheckOrigin allows all origins to connect (customize this based on security needs).
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}
