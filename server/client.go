package server

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	conn        *websocket.Conn
	id          string
	code        string
	readChannel chan []byte
	errChannel  chan error
	writeMutex  sync.Mutex
}

func newClient(id string, conn *websocket.Conn, code string) *client {
	return &client{
		conn:        conn,
		id:          id,
		code:        code,
		readChannel: make(chan []byte),
		errChannel:  make(chan error),
	}
}

func (c *client) handleRead() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.errChannel <- err
			return
		}
		c.readChannel <- message
	}
}

func (c *client) run(channel chan ClientCommand) {
	defer c.conn.Close()
	go c.handleRead()

	for {
		select {
		case msg := <-c.readChannel:
			fmt.Printf("Received message from client %s: %s\n", c.id, string(msg))
			channel <- ClientCommand{id: c.id, message: msg}

		case err := <-c.errChannel:
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			return

		case <-time.After(60 * time.Second):
			log.Printf("Client %s timed out due to inactivity", c.id)
			return
		}
	}
}

func (c *client) send(message []byte) error {
	c.writeMutex.Lock()
	defer c.writeMutex.Unlock()

	return c.conn.WriteMessage(websocket.TextMessage, message)
}
