package lobby

type EventJoined struct {
	Name string `json:"name"`
}

func (e EventJoined) Kind() string {
	return "joined"
}
