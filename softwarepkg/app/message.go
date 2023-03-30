package app

import (
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/code"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repo"
)

type MessageService interface {
	HandleNewPkg(CmdToHandleNewPkg) error
}

func NewMessageService(
	p repo.Repo,
	s message.SoftwarePkgProducer,
	c code.Code,
) *messageService {
	return &messageService{
		pr:       p,
		producer: s,
		code:     c,
	}
}

type messageService struct {
	pr       repo.Repo
	producer message.SoftwarePkgProducer
	code     code.Code
}

func (m *messageService) HandleNewPkg(cmd CmdToHandleNewPkg) error {
	if cmd.Platform != domain.PlatformGithub {
		return nil
	}

	url, err := m.pr.CreateRepo(cmd.PkgName)
	if err != nil {
		return err
	}

	e := domain.NewRepoCreatedEvent(cmd.PkgId, url)
	if err = m.producer.NotifyRepoCreatedResult(&e); err != nil {
		return err
	}

	e = domain.NewCodePushedEvent(cmd.PkgId)
	v := domain.NewSoftwarePkg(
		cmd.PkgName,
		domain.Importer{
			Name:  cmd.Importer,
			Email: cmd.ImporterEmail,
		},
		domain.SourceCode{
			SpecURL:   cmd.SpecURL,
			SrcRPMURL: cmd.SrcRPMURL,
		},
	)
	if err = m.code.Push(&v); err != nil {
		e.FailedReason = err.Error()
	}

	return m.producer.NotifyCodePushedResult(&e)
}
