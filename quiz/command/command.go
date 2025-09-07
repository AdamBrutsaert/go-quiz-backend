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
		var cmd Register
		if err := json.Unmarshal(base.Data, &cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	case commandKindStart:
		var cmd Start
		if err := json.Unmarshal(base.Data, &cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	default:
		return nil, errors.New("unknown command kind: " + base.Kind)
	}
}
