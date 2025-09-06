package lobby

import (
	"encoding/json"
	"errors"
)

type Command interface {
	Handle(id string, lobby *Lobby) error
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
	case commandRegisterKind:
		var event CommandRegister
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	case commandStartKind:
		var event CommandStart
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, errors.New("unknown command kind: " + base.Kind)
	}
}
