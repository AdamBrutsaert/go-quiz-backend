package lobby

import (
	"fmt"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type Lobby struct {
	notifier quiz.Notifier
	owner    string
	players  map[string]quiz.Player
}

func New(notifier quiz.Notifier) *Lobby {
	return &Lobby{
		notifier: notifier,
		owner:    "",
		players:  make(map[string]quiz.Player),
	}
}

func (l *Lobby) Handle(id string, message []byte) {
	event, err := deserializeCommand(message)
	if err != nil {
		fmt.Printf("error deserializing command: %v\n", err)
		return
	}

	fmt.Printf("Handling event: %T\n", event)
	err = event.Handle(id, l)
	if err != nil {
		fmt.Printf("error handling event: %v\n", err)
		return
	}
}
