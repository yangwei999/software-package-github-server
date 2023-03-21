package app

import (
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repo"
)

type MessageService interface {
	CreateRepo(CmdToCreateRepo) error
}

func NewMessageService(
	p repo.Repo,
	s message.SoftwarePkgProducer,
) *messageService {
	return &messageService{
		pr:       p,
		producer: s,
	}
}

type messageService struct {
	pr       repo.Repo
	producer message.SoftwarePkgProducer
}

func (m *messageService) CreateRepo(cmd CmdToCreateRepo) error {
	if cmd.Platform != domain.PlatformGithub {
		return nil
	}

	url, err := m.pr.CreateRepo(cmd.PkgName)
	if err != nil {
		return err
	}

	e := domain.NewRepoCreatedEvent(cmd.PkgId, url)
	return m.producer.NotifyRepoCreatedResult(&e)
}
