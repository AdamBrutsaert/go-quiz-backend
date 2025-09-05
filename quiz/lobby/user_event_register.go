package lobby

import "github.com/AdamBrutsaert/go-quiz-backend/quiz"

const userEventRegisterKind = "register"

type UserEventRegister struct {
	Name string `json:"name"`
}

func (e UserEventRegister) Kind() string {
	return userEventRegisterKind
}

func (e UserEventRegister) Handle(id string, lobby *Lobby) (quiz.Phase, error) {
	// If the user is already registered, do nothing
	if _, exists := lobby.players[id]; exists {
		return nil, nil
	}

	// If there is already an user with the same name, do nothing
	for _, player := range lobby.players {
		if player.Name == e.Name {
			return nil, nil
		}
	}

	// Register the user
	lobby.players[id] = quiz.Player{Name: e.Name}

	// If there is no owner yet, make this user the owner
	if lobby.owner == "" {
		lobby.owner = id
	}

	return nil, nil
}
