package lobby

import "github.com/AdamBrutsaert/go-quiz-backend/quiz"

type Lobby struct {
	players map[string]quiz.Player
	over    bool
}

func New() *Lobby {
	return &Lobby{
		players: make(map[string]quiz.Player),
		over:    false,
	}
}

func (g *Lobby) Players() map[string]quiz.Player {
	return g.players
}

func (g *Lobby) AddPlayer(name string) {
	g.players[name] = quiz.Player{Name: name}
}

func (g *Lobby) RemovePlayer(name string) {
	delete(g.players, name)
}

func (g *Lobby) Start() {
	g.over = true
}

func (g *Lobby) Over() bool {
	return g.over
}
