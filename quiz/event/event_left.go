package event

type Left struct {
	Name string `json:"name"`
}

func (Left) Kind() string {
	return "left"
}
