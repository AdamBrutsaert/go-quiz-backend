package command

import "github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"

const startCommandKind = "start"

type Start struct{}

func (e Start) Kind() string {
	return startCommandKind
}

func (e Start) Handle(g *lobby.Lobby) error {
	g.Start()
	return nil
}
