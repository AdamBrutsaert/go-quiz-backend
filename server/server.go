package server

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	quizzes  map[string]*runner
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
		quizzes: make(map[string]*runner),
	}
}

func (s *Server) createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealthCheck)
	mux.HandleFunc("/lobby", s.handleCreateLobby)
	mux.HandleFunc("/ws", s.handleWebSocket)
	return mux
}

func (s *Server) createServer() *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        s.createMux(),
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *Server) generateQuizCode() string {
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
		if _, exists := s.quizzes[code]; !exists {
			return code
		}
	}
}

func (s *Server) createQuiz() string {
	code := s.generateQuizCode()
	runner := newRunner()
	s.quizzes[code] = runner
	go runner.run()
	return code
}

func (s *Server) Run() error {
	return s.createServer().ListenAndServe()
}
