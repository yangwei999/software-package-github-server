package message

type EventMessage interface {
	Message() ([]byte, error)
}

type SoftwarePkgProducer interface {
	NotifyCodePushedResult(EventMessage) error
}
