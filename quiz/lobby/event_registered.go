package lobby

type EventRegistered struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e EventRegistered) Kind() string {
	return "registered"
}
