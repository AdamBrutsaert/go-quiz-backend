package main

import (
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/game"
	gameCommand "github.com/AdamBrutsaert/go-quiz-backend/quiz/game/command"
	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
	lobbyCommand "github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby/command"
)

func receivedLobbyCommands(yield func(lobbyCommand.Command) bool) {
	cmds := []lobbyCommand.Command{
		lobbyCommand.Register{Name: "Alice"},
		lobbyCommand.Register{Name: "Bob"},
		lobbyCommand.Start{},
	}

	for _, cmd := range cmds {
		if !yield(cmd) {
			return
		}
	}
}

func receivedGameCommands(yield func(gameCommand.Command) bool) {
	cmds := []gameCommand.Command{}

	for _, cmd := range cmds {
		if !yield(cmd) {
			return
		}
	}
}

func main() {
	lobby := lobby.New()
	for !lobby.Over() {
		for cmd := range receivedLobbyCommands {
			if err := cmd.Handle(&lobby); err != nil {
				panic(err)
			}
		}
	}

	game := game.New(lobby)
	for !game.Over() {
		for cmd := range receivedGameCommands {
			if err := cmd.Handle(&game); err != nil {
				panic(err)
			}
		}
	}
}
