package command

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

func disconnectPlayer(eventHandler state.EventHandler, players map[string]quiz.Player, owner *string, clientID string) {
	player, exists := players[clientID]
	if !exists {
		return
	}

	eventHandler.NotifyAllClients(event.Left{Name: player.Name})
	log.Printf("[%s][%s] Left", clientID, player.Name)

	delete(players, clientID)

	// If the owner left, assign a new owner
	if owner != nil && *owner == clientID {
		*owner = ""
		for id, player := range players {
			*owner = id
			eventHandler.NotifyAllClients(event.OwnerChanged{Name: player.Name})
			log.Printf("[%s][%s] Became the owner", id, player.Name)
			break
		}
	}
}

type Disconnect struct{}

func (e Disconnect) ExecuteLobby(lobby *state.Lobby, clientID string) {
	disconnectPlayer(lobby.EventHandler, lobby.Players, &lobby.Owner, clientID)
}

func (e Disconnect) ExecuteGame(game *state.Game, clientID string) {
	disconnectPlayer(game.EventHandler, game.Players, &game.Owner, clientID)
}
