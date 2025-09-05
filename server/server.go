package server

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	lobbies  map[string]lobby.Lobby
}

func New() *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		lobbies: make(map[string]lobby.Lobby),
	}
}

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

func (s *Server) generateLobbyCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 5)
	for {
		if _, err := rand.Read(b); err != nil {
			log.Printf("Error generating random bytes: %v", err)
			continue
		}

		for i := range b {
			b[i] = charset[b[i]%byte(len(charset))]
		}

		code := string(b)
		if _, exists := s.lobbies[code]; !exists {
			return code
		}
	}
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

func (s *Server) createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleWebSocket)
	mux.HandleFunc("/lobby", s.handleCreateLobby)
	return mux
}

func (s *Server) createServer() *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        s.createMux(),
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *Server) Run() error {
	return s.createServer().ListenAndServe()
}
