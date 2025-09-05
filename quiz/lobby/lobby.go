package lobby

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type Lobby struct {
	owner   string
	players map[string]quiz.Player
}

func New() *Lobby {
	return &Lobby{
		owner:   "",
		players: make(map[string]quiz.Player),
	}
}

func (l *Lobby) Handle(id string, message []byte) (quiz.Phase, error) {
	event, err := deserializeUserEvent(message)
	if err != nil {
		return nil, err
	}
	return event.Handle(id, l)
}
