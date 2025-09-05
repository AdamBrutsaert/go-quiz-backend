package quiz

type Phase interface {
	Handle(id string, message []byte) (Phase, error)
}
