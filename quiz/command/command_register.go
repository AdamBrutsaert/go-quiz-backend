package command

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

type Register struct {
	Name string `json:"name"`
}

func (e Register) ExecuteLobby(lobby *state.Lobby, clientID string) {
	if e.Name == "" {
		lobby.EventHandler.NotifyClient(clientID, event.ErrInvalidName)
		return
	}

	if _, exists := lobby.Players[clientID]; exists {
		lobby.EventHandler.NotifyClient(clientID, event.ErrAlreadyRegistered)
		return
	}

	for _, player := range lobby.Players {
		if player.Name == e.Name {
			lobby.EventHandler.NotifyClient(clientID, event.ErrNameAlreadyTaken)
			return
		}
	}

	lobby.Players[clientID] = quiz.Player{
		Name:  e.Name,
		Score: 0,
	}
	log.Printf("[%s][%s] Registered", clientID, e.Name)

	lobby.EventHandler.NotifyClient(clientID, event.Registered{ID: clientID, Name: e.Name})
	lobby.EventHandler.NotifyAllClientsExcept(clientID, event.Joined{Name: e.Name})

	if lobby.Owner == "" {
		lobby.Owner = clientID
		lobby.EventHandler.NotifyAllClients(event.OwnerChanged{Name: e.Name})
		log.Printf("[%s][%s] Became the owner", clientID, e.Name)
	}
}

func (e Register) ExecuteGame(game *state.Game, clientID string) {
	game.EventHandler.NotifyClient(clientID, event.ErrInvalidCommand)
}
