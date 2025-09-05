package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocketUpgrade(t *testing.T) {
	s := New()
	ts := httptest.NewServer(s.createMux())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	u.Scheme = "ws"
	u.Path = "/ws"

	ws, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("WebSocket upgrade failed: %v", err)
	}
	defer ws.Close()

	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Errorf("Expected status %d, got %d", http.StatusSwitchingProtocols, resp.StatusCode)
	}
}

func TestCreateLobby(t *testing.T) {
	s := New()
	req := httptest.NewRequest(http.MethodPost, "/lobby", nil)
	w := httptest.NewRecorder()

	s.handleCreateLobby(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	code, exists := response["code"]
	if !exists {
		t.Error("Response should contain 'code' field")
	}

	if len(code) != 5 {
		t.Errorf("Expected code length 5, got %d", len(code))
	}

	// Verify the lobby was stored in the server
	if _, exists := s.lobbies[code]; !exists {
		t.Error("Lobby should be stored in server")
	}
}

func TestCreateLobbyMethodNotAllowed(t *testing.T) {
	s := New()
	req := httptest.NewRequest(http.MethodGet, "/lobby", nil)
	w := httptest.NewRecorder()

	s.handleCreateLobby(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestCreateLobbyUniqueCodes(t *testing.T) {
	s := New()
	codes := make(map[string]bool)

	// Create multiple lobbies and verify all codes are unique
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest(http.MethodPost, "/lobby", nil)
		w := httptest.NewRecorder()

		s.handleCreateLobby(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]string
		err := json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		code := response["code"]
		if codes[code] {
			t.Errorf("Duplicate code generated: %s", code)
		}
		codes[code] = true
	}
}
