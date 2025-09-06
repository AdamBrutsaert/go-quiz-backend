package lobby

import (
	"errors"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/game"
)

const commandStartKind = "start"

type CommandStart struct{}

func (e CommandStart) Handle(id string, lobby *Lobby) error {
	// Only the owner can start the game
	if id != lobby.owner {
		return errors.New("only the owner can start the game")
	}

	lobby.notifier.NotifyPhase(game.New(lobby.notifier, lobby.players))
	return nil
}
