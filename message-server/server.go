package messageserver

import (
	"context"
	"encoding/json"

	"github.com/opensourceways/software-package-github-server/mq"
	"github.com/opensourceways/software-package-github-server/softwarepkg/app"
)

func Init(s app.MessageService, c Config) *MessageServer {
	return &MessageServer{
		cfg:     c,
		service: s,
	}
}

type MessageServer struct {
	cfg     Config
	service app.MessageService
}

func (m *MessageServer) Run(ctx context.Context) error {
	if err := m.subscribe(); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (m *MessageServer) subscribe() error {
	h := map[string]mq.Handler{
		m.cfg.Topics.SoftwarePkgInitialized: m.handlePkgInitialized,
	}

	return mq.Subscriber().Subscribe(m.cfg.Group, h)
}

func (m *MessageServer) handlePkgInitialized(data []byte) error {
	msg := new(msgToHandlePkgInitialized)

	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	return m.service.HandlePkgInitialized(*msg)
}
