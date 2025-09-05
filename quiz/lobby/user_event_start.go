package lobby

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/game"
)

const userEventStartKind = "start"

type UserEventStart struct{}

func (e UserEventStart) Kind() string {
	return userEventStartKind
}

func (e UserEventStart) Handle(id string, lobby *Lobby) (quiz.Phase, error) {
	// Only the owner can start the game
	if id != lobby.owner {
		return nil, nil
	}

	return game.New(lobby.players), nil
}
