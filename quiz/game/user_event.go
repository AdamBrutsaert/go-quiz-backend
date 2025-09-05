package game

import (
	"encoding/json"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type UserEvent interface {
	Kind() string
	Handle(id string, lobby *Game) (quiz.Phase, error)
}

func deserializeUserEvent(data []byte) (UserEvent, error) {
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
