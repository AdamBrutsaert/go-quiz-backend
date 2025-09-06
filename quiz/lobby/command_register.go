package lobby

import (
	"errors"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

const commandRegisterKind = "register"

type CommandRegister struct {
	Name string `json:"name"`
}

func (e CommandRegister) Handle(id string, lobby *Lobby) error {
	// Validate name
	if e.Name == "" {
		return errors.New("name cannot be empty")
	}

	// If the user is already registered, do nothing
	if _, exists := lobby.players[id]; exists {
		return errors.New("user already registered")
	}

	// If there is already an user with the same name, do nothing
	for _, player := range lobby.players {
		if player.Name == e.Name {
			return errors.New("name already taken")
		}
	}

	// Register the user
	lobby.players[id] = quiz.Player{Name: e.Name}

	// If there is no owner yet, make this user the owner
	if lobby.owner == "" {
		lobby.owner = id
	}

	lobby.notifyOne(id, EventRegistered{ID: id, Name: e.Name})
	lobby.notifyAll(EventJoined{Name: e.Name})

	return nil
}
