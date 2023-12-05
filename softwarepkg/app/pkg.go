package app

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/code"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/useradapter"
)

type PkgService interface {
	HandlePushCode(*CmdToHandlePushCode) error
}

func NewPkgService(
	c code.Code,
	prod message.SoftwarePkgProducer,
	u useradapter.UserAdapter,
) *pkgService {
	return &pkgService{
		code:     c,
		producer: prod,
		user:     u,
	}
}

type pkgService struct {
	code     code.Code
	producer message.SoftwarePkgProducer
	user     useradapter.UserAdapter
}

func (p *pkgService) HandlePushCode(cmd *CmdToHandlePushCode) error {
	if cmd.Platform != domain.PlatformGithub {
		return nil
	}

	importerEmail, err := p.user.GetEmail(cmd.Importer)
	if err != nil {
		return err
	}

	if !p.code.CheckRepoCreated(cmd.PkgName) {
		return fmt.Errorf("repo %s has not been created", cmd.PkgName)
	}

	pushCode := cmd.toPushCode(importerEmail)
	repoUrl, err := p.code.Push(&pushCode)
	if err != nil {
		logrus.Errorf("pkgId %s push code err: %s", pushCode.PkgId, err.Error())

		return err
	}

	e := domain.NewCodePushedEvent(pushCode.PkgId, repoUrl)
	return p.producer.NotifyCodePushedResult(&e)
}
