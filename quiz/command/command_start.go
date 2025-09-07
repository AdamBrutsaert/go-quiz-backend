package command

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

type Start struct{}

func (e Start) ExecuteLobby(lobby *state.Lobby, clientID string) error {
	if lobby.Owner != clientID {
		return quiz.ErrNotOwner
	}

	game := state.NewGame(lobby)
	lobby.EventHandler.NotifyNewState(game)

	log.Printf("[%s] Started the game", clientID)
	return nil
}

func (e Start) ExecuteGame(game *state.Game, clientID string) error {
	return quiz.ErrInvalidCommand
}
