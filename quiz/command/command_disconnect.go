package command

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

type Disconnect struct{}

func (e Disconnect) ExecuteLobby(lobby *state.Lobby, clientID string) error {
	player, exists := lobby.Players[clientID]
	if !exists {
		return nil
	}

	lobby.EventHandler.NotifyAllClients(event.Left{Name: player.Name})
	log.Printf("[%s][%s] Left", clientID, player.Name)

	delete(lobby.Players, clientID)

	// If the owner left, assign a new owner
	if lobby.Owner == clientID {
		lobby.Owner = ""
		for id, player := range lobby.Players {
			lobby.Owner = id
			lobby.EventHandler.NotifyAllClients(event.OwnerChanged{Name: player.Name})
			log.Printf("[%s][%s] Became the owner", id, player.Name)
			break
		}
	}

	return nil
}

func (e Disconnect) ExecuteGame(game *state.Game, clientID string) error {
	player, exists := game.Players[clientID]
	if !exists {
		return nil
	}

	game.EventHandler.NotifyAllClients(event.Left{Name: player.Name})
	log.Printf("[%s][%s] Left", clientID, player.Name)

	delete(game.Players, clientID)

	// If the owner left, assign a new owner
	if game.Owner == clientID {
		game.Owner = ""
		for id, player := range game.Players {
			game.Owner = id
			game.EventHandler.NotifyAllClients(event.OwnerChanged{Name: player.Name})
			log.Printf("[%s][%s] Became the owner", id, player.Name)
			break
		}
	}

	return nil
}
