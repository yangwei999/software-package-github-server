package messageimpl

import (
	"github.com/opensourceways/software-package-github-server/mq"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
)

func NewMessageImpl(c Config) *MessageImpl {
	return &MessageImpl{
		cfg: c,
	}
}

type MessageImpl struct {
	cfg Config
}

func (m *MessageImpl) NotifyRepoCreatedResult(msg message.EventMessage) error {
	return send(m.cfg.TopicsToNotify.CreatedRepo, msg)
}

func send(topic string, v message.EventMessage) error {
	body, err := v.Message()
	if err != nil {
		return err
	}

	return mq.Subscriber().Publish(topic, body)
}
