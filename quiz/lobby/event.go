package lobby

import "encoding/json"

type Event interface {
	Kind() string
}

func serializeEvent(event Event) ([]byte, error) {
	var base struct {
		Kind string `json:"kind"`
		Data Event  `json:"data"`
	}

	base.Kind = event.Kind()
	base.Data = event

	return json.Marshal(base)
}

func (l *Lobby) notifyOne(id string, event Event) error {
	data, err := serializeEvent(event)
	if err != nil {
		return err
	}
	l.notifier.NotifyOne(id, data)
	return nil
}

func (l *Lobby) notifyAll(event Event) error {
	data, err := serializeEvent(event)
	if err != nil {
		return err
	}
	l.notifier.NotifyAll(data)
	return nil
}
