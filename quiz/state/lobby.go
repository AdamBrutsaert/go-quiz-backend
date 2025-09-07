package state

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type Lobby struct {
	EventHandler EventHandler
	Owner        string
	Players      map[string]quiz.Player
}

func NewLobby(eventHandler EventHandler) *Lobby {
	return &Lobby{
		EventHandler: eventHandler,
		Owner:        "",
		Players:      make(map[string]quiz.Player),
	}
}

func (l *Lobby) Start() error {
	return nil
}

func (l *Lobby) Apply(command Command, clientID string) {
	command.ExecuteLobby(l, clientID)
}
