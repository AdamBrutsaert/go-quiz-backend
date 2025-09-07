package event

type Registered struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Registered) Kind() string {
	return "registered"
}
