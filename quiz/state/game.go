package state

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
)

type Game struct {
	EventHandler EventHandler
	Owner        string
	Players      map[string]quiz.Player
}

func NewGame(lobby *Lobby) *Game {
	return &Game{
		EventHandler: lobby.EventHandler,
		Owner:        lobby.Owner,
		Players:      lobby.Players,
	}
}

func (g *Game) Start() error {
	g.EventHandler.NotifyAllClients(event.Start{})
	return nil
}

func (g *Game) Apply(command Command, clientID string) {
	command.ExecuteGame(g, clientID)
}
