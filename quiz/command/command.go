package command

import (
	"encoding/json"
	"errors"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz/state"
)

const (
	commandKindRegister = "register"
	commandKindStart    = "start"
)

func Deserialize(data []byte) (state.Command, error) {
	var base struct {
		Kind string          `json:"kind"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Kind {
	case commandKindRegister:
		var event Register
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	case commandKindStart:
		var event Start
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, errors.New("unknown command kind: " + base.Kind)
	}
}
