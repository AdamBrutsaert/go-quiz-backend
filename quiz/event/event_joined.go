package event

type Joined struct {
	Name string `json:"name"`
}

func (Joined) Kind() string {
	return "joined"
}
