package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	quizzes  map[string]*Runner
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
		quizzes: make(map[string]*Runner),
	}
}

func (s *Server) createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleWebSocket)
	mux.HandleFunc("/lobby", s.handleCreateLobby)
	mux.HandleFunc("/health", s.handleHealthCheck)
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

func (s *Server) newQuiz() string {
	code := s.generateQuizCode()
	runner := newRunner()
	s.quizzes[code] = runner
	go runner.run()
	return code
}
