package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	conn        *websocket.Conn
	id          string
	code        string
	readChannel chan []byte
	errChannel  chan error
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

func (c *client) run(server *Server) {
	defer c.conn.Close()
	go c.handleRead()

	for {
		select {
		case msg := <-c.readChannel:
			quiz := server.Quiz(c.code)
			if quiz == nil {
				log.Printf("Quiz not found for code: %s", c.code)
				return
			}

			phase, err := quiz.Handle(c.id, msg)
			if err != nil {
				log.Printf("Error handling message: %v", err)
				continue
			}

			if phase != nil {
				server.SetQuiz(c.code, phase)
			}

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
	return c.conn.WriteMessage(websocket.TextMessage, message)
}
