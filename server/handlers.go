package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
	"github.com/gorilla/websocket"
)

func (s *Server) handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	for {
		// Read message from WebSocket
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		log.Printf("Received message of type %d: %s", messageType, message)
	}
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not open websocket connection: %v", err)
		return
	}

	go s.handleConnection(conn)
}

func (s *Server) handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := s.generateLobbyCode()

	newLobby := lobby.New()
	s.lobbies[code] = newLobby

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"code": code}
	json.NewEncoder(w).Encode(response)
}
