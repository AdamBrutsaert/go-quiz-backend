package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
	"github.com/gorilla/websocket"
)

type Quiz struct {
	phaseHandleMutex sync.Mutex
	phaseUpdateMutex sync.Mutex
	phase            quiz.Phase
	clients          map[string]*client
}

func (q *Quiz) NotifyOne(id string, message []byte) {
	if client, ok := q.clients[id]; ok {
		err := client.send(message)
		if err != nil {
			fmt.Println("Error sending message to client:", err)
		}
	}
}

func (q *Quiz) NotifyAll(message []byte) {
	for _, client := range q.clients {
		err := client.send(message)
		if err != nil {
			fmt.Println("Error sending message to client:", err)
		}
	}
}

func (q *Quiz) NotifyPhase(phase quiz.Phase) {
	handleLocked := q.phaseHandleMutex.TryLock()
	q.phaseUpdateMutex.Lock()

	q.phase = phase

	q.phaseUpdateMutex.Unlock()
	if handleLocked {
		q.phaseHandleMutex.Unlock()
	}
}

type Server struct {
	upgrader websocket.Upgrader
	quizzes  map[string]*Quiz
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
		quizzes: make(map[string]*Quiz),
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

func (s *Server) createQuiz() string {
	code := s.generateQuizCode()
	quiz := &Quiz{
		phase:   nil,
		clients: make(map[string]*client),
	}
	quiz.phase = lobby.New(quiz)
	s.quizzes[code] = quiz
	return code
}
