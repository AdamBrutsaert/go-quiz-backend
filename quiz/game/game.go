package game

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type Game struct {
	players map[string]quiz.Player
}

func New(players map[string]quiz.Player) *Game {
	return &Game{
		players: players,
	}
}

func (g *Game) Handle(id string, message []byte) (quiz.Phase, error) {
	event, err := deserializeUserEvent(message)
	if err != nil {
		return nil, err
	}
	return event.Handle(id, g)
}
