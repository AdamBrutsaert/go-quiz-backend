package server

import (
	"log"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/command"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/event"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
	"github.com/google/uuid"
)

type clientCommand struct {
	id  string
	cmd state.Command
}

type runner struct {
	state   state.State
	clients map[string]*client

	commandsChannel   chan clientCommand
	disconnectChannel chan string
	stateChannel      chan state.State
}

func newRunner() *runner {
	runner := &runner{
		clients:           make(map[string]*client),
		commandsChannel:   make(chan clientCommand, 100),
		disconnectChannel: make(chan string, 10),
		stateChannel:      make(chan state.State, 1),
	}
	runner.state = state.NewLobby(runner)
	return runner
}

func (r *runner) NotifyNewState(s state.State) {
	r.stateChannel <- s
}

func (r *runner) NotifyClient(clientID string, ev event.Event) {
	client, exists := r.clients[clientID]
	if !exists {
		log.Printf("Client %s not found for notification\n", clientID)
		return
	}

	serialized, err := event.Serialize(ev)
	if err != nil {
		log.Printf("Error serializing event %v: %v\n", ev, err)
		return
	}

	client.write(serialized)
}

func (r *runner) NotifyAllClients(ev event.Event) {
	serialized, err := event.Serialize(ev)
	if err != nil {
		log.Printf("Error serializing event %v: %v\n", ev, err)
		return
	}

	for _, client := range r.clients {
		client.write(serialized)
	}
}

func (r *runner) NotifyAllClientsExcept(clientID string, ev event.Event) {
	serialized, err := event.Serialize(ev)
	if err != nil {
		log.Printf("Error serializing event %v: %v\n", ev, err)
		return
	}

	for _, client := range r.clients {
		if client.id == clientID {
			continue
		}
		client.write(serialized)
	}
}

func (r *runner) generateClientID() string {
	return uuid.New().String()
}

func (r *runner) run() {
	r.state.Start()

	for {
		select {
		case cmd := <-r.commandsChannel:
			r.state.Apply(cmd.cmd, cmd.id)
		case clientID := <-r.disconnectChannel:
			delete(r.clients, clientID)
			r.state.Apply(command.Disconnect{}, clientID)
		case state := <-r.stateChannel:
			r.state = state
			r.state.Start()
		}
	}
}
