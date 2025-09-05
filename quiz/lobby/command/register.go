package command

import "github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"

const registerCommandKind = "register"

type Register struct {
	Name string `json:"name"`
}

func (e Register) Kind() string {
	return registerCommandKind
}

func (e Register) Handle(g *lobby.Lobby) error {
	g.AddPlayer(e.Name)
	return nil
}
