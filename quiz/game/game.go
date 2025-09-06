package game

import (
	"fmt"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type Game struct {
	notifier quiz.Notifier
	players  map[string]quiz.Player
}

func New(notifier quiz.Notifier, players map[string]quiz.Player) *Game {
	return &Game{
		notifier: notifier,
		players:  players,
	}
}

func (l *Game) Handle(id string, message []byte) {
	event, err := deserializeCommand(message)
	if err != nil {
		fmt.Printf("error deserializing command: %v\n", err)
		return
	}

	err = event.Handle(id, l)
	if err != nil {
		fmt.Printf("error handling event: %v\n", err)
		return
	}
}
