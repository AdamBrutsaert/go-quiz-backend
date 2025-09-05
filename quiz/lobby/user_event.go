package lobby

import (
	"encoding/json"

	"github.com/AdamBrutsaert/go-quiz-backend/quiz"
)

type UserEvent interface {
	Kind() string
	Handle(id string, lobby *Lobby) (quiz.Phase, error)
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
	case userEventRegisterKind:
		var event UserEventRegister
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	case userEventStartKind:
		var event UserEventStart
		if err := json.Unmarshal(base.Data, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, nil
	}
}
