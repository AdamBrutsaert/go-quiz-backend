package event

import "encoding/json"

type Event interface {
	Kind() string
}

func Serialize(event Event) ([]byte, error) {
	var base struct {
		Kind string `json:"kind"`
		Data Event  `json:"data"`
	}

	base.Kind = event.Kind()
	base.Data = event

	return json.Marshal(base)
}
