package server

import (
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
