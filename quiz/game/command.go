package game

import (
	"encoding/json"
	"errors"
)

type Command interface {
	Handle(id string, lobby *Game) error
}

func deserializeCommand(data []byte) (Command, error) {
	var base struct {
		Kind string          `json:"kind"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Kind {
	default:
		return nil, errors.New("unknown command kind: " + base.Kind)
	}
}
