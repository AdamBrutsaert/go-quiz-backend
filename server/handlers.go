package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Get the code query parameter
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	// Verify the quiz exists
	quiz, exists := s.quizzes[code]
	if !exists {
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}

	// Upgrade the connection to WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not open websocket connection: %v", err)
		return
	}

	client := newClient(generateClientID(), conn, code)
	quiz.clients[client.id] = client

	go client.run(quiz.commandsChannel)
}

func (s *Server) handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := s.newQuiz()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"code": code}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
