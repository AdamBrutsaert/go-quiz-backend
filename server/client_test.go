package server

import (
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestClientRunTimeout(t *testing.T) {
	// Create a server for WebSocket connection
	server := New()
	ts := httptest.NewServer(server.createMux())
	defer ts.Close()

	// Create a lobby first
	testCode := server.CreateLobby()

	u, _ := url.Parse(ts.URL)
	u.Scheme = "ws"
	u.Path = "/ws"
	u.RawQuery = "code=" + testCode

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("Failed to create WebSocket connection: %v", err)
	}
	defer conn.Close()

	// Create a client
	c := &client{
		conn: conn,
		id:   generateClientID(),
		code: testCode,
	}

	// Test that run method can be called without hanging
	done := make(chan bool, 1)
	go func() {
		c.run(server)
		done <- true
	}()

	// Close the connection to trigger the run method to exit
	conn.Close()

	// Wait for the run method to complete or timeout
	select {
	case <-done:
		// Success - run method completed
	case <-time.After(2 * time.Second):
		t.Error("Client run method did not complete within timeout")
	}
}
