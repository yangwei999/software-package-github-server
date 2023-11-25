package messageserver

import (
	"context"
	"encoding/json"

	kafka "github.com/opensourceways/kafka-lib/agent"

	"github.com/opensourceways/software-package-github-server/softwarepkg/app"
)

func Init(s app.PkgService, c Config) *MessageServer {
	return &MessageServer{
		cfg:     c,
		service: s,
	}
}

type MessageServer struct {
	cfg     Config
	service app.PkgService
}

func (m *MessageServer) Run(ctx context.Context) error {
	err := kafka.Subscribe(m.cfg.Group, m.handlePushCode, []string{m.cfg.Topics.PushCode})
	if err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (m *MessageServer) handlePushCode(payload []byte, header map[string]string) error {
	msg := new(msgToHandlePushCode)

	if err := json.Unmarshal(payload, msg); err != nil {
		return err
	}

	return m.service.HandlePushCode(msg)
}
