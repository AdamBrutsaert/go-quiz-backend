package quiz

type Phase interface {
	Handle(id string, message []byte)
}

type Notifier interface {
	NotifyOne(id string, message []byte)
	NotifyAll(message []byte)
	NotifyPhase(phase Phase)
}
