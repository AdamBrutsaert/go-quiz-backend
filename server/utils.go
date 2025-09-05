package server

import (
	"crypto/rand"
	"log"

	"github.com/google/uuid"
)

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
		if _, exists := s.quizes[code]; !exists {
			return code
		}
	}
}

func generateClientID() string {
	return uuid.New().String()
}
