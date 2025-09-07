package event

type Start struct {
}

func (Start) Kind() string {
	return "start"
}
