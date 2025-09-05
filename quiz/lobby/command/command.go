package command

import (
	"encoding/json"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/lobby"
)

type Command interface {
	Kind() string
	Handle(g *lobby.Lobby) error
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
	case registerCommandKind:
		var event Register
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	case startCommandKind:
		var event Start
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, nil
	}
}
