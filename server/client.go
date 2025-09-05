package server

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
	id   string
	code string
}

func (c *client) run(server *Server) {
	defer c.conn.Close()

	for {
		// Read message from WebSocket
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		log.Printf("Received message of type %d: %s", messageType, message)
	}
}
