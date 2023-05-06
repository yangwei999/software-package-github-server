package app

import (
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repository"
)

type MessageService interface {
	HandleNewPkg(CmdToHandleNewPkg) error
}

func NewMessageService(r repository.SoftwarePkg) *messageService {
	return &messageService{
		repository: r,
	}
}

type messageService struct {
	repository repository.SoftwarePkg
}

func (m *messageService) HandleNewPkg(cmd CmdToHandleNewPkg) error {
	if cmd.Platform != domain.PlatformGithub {
		return nil
	}

	pkg := domain.NewSoftwarePkg(
		domain.SoftwarePkgBasic{
			Id:   cmd.PkgId,
			Name: cmd.PkgName,
		},
		domain.Importer{
			Name:  cmd.Importer,
			Email: cmd.ImporterEmail,
		},
		domain.SourceCode{
			SpecURL:   cmd.SpecURL,
			SrcRPMURL: cmd.SrcRPMURL,
		},
		cmd.CIPRNum,
	)

	return m.repository.Add(&pkg)
}
