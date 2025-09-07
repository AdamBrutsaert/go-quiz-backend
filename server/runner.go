package server

import (
	"fmt"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
)

type ClientCommand struct {
	id      string
	message []byte
}

type Runner struct {
	phase           quiz.Phase
	commandsChannel chan ClientCommand
	phaseChannel    chan quiz.Phase
	clients         map[string]*client
}

func newRunner() *Runner {
	runner := &Runner{
		commandsChannel: make(chan ClientCommand, 100),
		phaseChannel:    make(chan quiz.Phase, 1),
		clients:         make(map[string]*client),
	}

	runner.phase = lobby.New(runner)
	return runner
}

func (r *Runner) run() {
	for {
		select {
		case cmd := <-r.commandsChannel:
			r.phase.Handle(cmd.id, cmd.message)
		case phase := <-r.phaseChannel:
			r.phase = phase
		}
	}
}

func (r *Runner) NotifyPhase(phase quiz.Phase) {
	r.phaseChannel <- phase
}

func (r *Runner) NotifyOne(id string, message []byte) {
	if client, ok := r.clients[id]; ok {
		err := client.send(message)
		if err != nil {
			fmt.Printf("Error sending message to client %s: %v\n", id, err)
		}
	}
}

func (r *Runner) NotifyAll(message []byte) {
	for _, client := range r.clients {
		err := client.send(message)
		if err != nil {
			fmt.Printf("Error sending message to client %s: %v\n", client.id, err)
		}
	}
}
