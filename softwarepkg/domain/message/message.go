package message

type EventMessage interface {
	Message() ([]byte, error)
}

type SoftwarePkgProducer interface {
	NotifyRepoCreatedResult(EventMessage) error
	NotifyCodePushedResult(EventMessage) error
}
