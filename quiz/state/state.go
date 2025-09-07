package state

import "github.com/AdamBrutsaert/go-quiz-backend/quiz/event"

type Command interface {
	ExecuteGame(game *Game, clientID string)
	ExecuteLobby(lobby *Lobby, clientID string)
}

type EventHandler interface {
	NotifyNewState(state State)
	NotifyClient(clientID string, event event.Event)
	NotifyAllClients(event event.Event)
	NotifyAllClientsExcept(clientID string, event event.Event)
}

type State interface {
	Start() error
	Apply(command Command, clientID string)
}
