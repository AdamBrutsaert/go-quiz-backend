package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
)

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Get the code query parameter
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	// Verify the lobby exists
	if _, exists := s.quizes[code]; !exists {
		http.Error(w, "Lobby not found", http.StatusNotFound)
		return
	}

	// Upgrade the connection to WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not open websocket connection: %v", err)
		return
	}

	// Create a new client with a unique ID and run it
	go newClient(generateClientID(), conn, code).run(s)
}

func (s *Server) handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := s.generateQuizCode()
	s.SetQuiz(code, lobby.New())

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"code": code}
	json.NewEncoder(w).Encode(response)
}
