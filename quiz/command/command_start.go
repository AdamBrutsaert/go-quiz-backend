package command

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

type Start struct{}

func (e Start) ExecuteLobby(lobby *state.Lobby, clientID string) {
	if lobby.Owner != clientID {
		lobby.EventHandler.NotifyClient(clientID, event.ErrNotOwner)
		return
	}

	game := state.NewGame(lobby)
	lobby.EventHandler.NotifyNewState(game)

	log.Printf("[%s] Started the game", clientID)
}

func (e Start) ExecuteGame(game *state.Game, clientID string) {
	game.EventHandler.NotifyClient(clientID, event.ErrInvalidCommand)
}
