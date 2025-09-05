package game

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
)

type Game struct {
	players map[string]quiz.Player
}

func New(lobby lobby.Lobby) Game {
	players := make(map[string]quiz.Player)
	for k, v := range lobby.Players() {
		players[k] = v
	}

	return Game{
		players: players,
	}
}

func (g Game) Over() bool {
	return false
}
