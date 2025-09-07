package event

type OwnerChanged struct {
	Name string `json:"name"`
}

func (OwnerChanged) Kind() string {
	return "owner_changed"
}
