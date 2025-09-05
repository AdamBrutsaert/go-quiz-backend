package command

import (
	"encoding/json"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/game"
)

type Command interface {
	Kind() string
	Handle(g *game.Game) error
}

func Serialize(command Command) ([]byte, error) {
	var base struct {
		Kind string  `json:"kind"`
		Data Command `json:"data"`
	}

	base.Kind = command.Kind()
	base.Data = command

	return json.Marshal(base)
}

func Deserialize(data []byte) (Command, error) {
	var base struct {
		Kind string          `json:"kind"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Kind {
	default:
		return nil, nil
	}
}
