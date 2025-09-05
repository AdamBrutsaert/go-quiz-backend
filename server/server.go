package server

import (
	"net/http"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	quizes   map[string]quiz.Phase
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
		quizes: make(map[string]quiz.Phase),
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

func (s *Server) SetQuiz(code string, phase quiz.Phase) {
	s.quizes[code] = phase
}

func (s *Server) Quiz(code string) quiz.Phase {
	return s.quizes[code]
}
