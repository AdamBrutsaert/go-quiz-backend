package server

import (
	"log"
	"time"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/command"
	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
	id   string

	readChannel  chan []byte
	writeChannel chan []byte
	errChannel   chan error
}

func newClient(conn *websocket.Conn, id string) *client {
	return &client{
		conn:         conn,
		id:           id,
		readChannel:  make(chan []byte, 10),
		writeChannel: make(chan []byte, 10),
		errChannel:   make(chan error),
	}
}

func (c *client) readPump() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.errChannel <- err
			return
		}
		c.readChannel <- message
	}
}

func (c *client) write(message []byte) {
	c.writeChannel <- message
}

func (c *client) run(commandsChannel chan clientCommand, disconnectChannel chan string) {
	defer c.conn.Close()
	go c.readPump() // will be cleaned up by the defer above

outer:
	for {
		select {
		case msg := <-c.readChannel:
			cmd, err := command.Deserialize(msg)
			if err != nil {
				log.Printf("[%s] Error deserializing command: %v\n", c.id, err)
				continue
			}

			commandsChannel <- clientCommand{id: c.id, cmd: cmd}

		case msg := <-c.writeChannel:
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("[%s] Error writing message: %v\n", c.id, err)
				break outer
			}

		case err := <-c.errChannel:
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[%s] WebSocket error: %v", c.id, err)
			}
			break outer

		case <-time.After(60 * time.Second):
			log.Printf("[%s] Timed out due to inactivity", c.id)
			break outer
		}
	}

	log.Printf("[%s] Connection closed", c.id)
	disconnectChannel <- c.id
}
