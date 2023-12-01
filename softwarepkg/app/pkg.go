package app

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/code"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
)

type PkgService interface {
	HandlePushCode(pushCode *domain.PushCode) error
}

func NewPkgService(c code.Code, prod message.SoftwarePkgProducer,
) *pkgService {
	return &pkgService{
		code:     c,
		producer: prod,
	}
}

type pkgService struct {
	code     code.Code
	producer message.SoftwarePkgProducer
}

func (p *pkgService) HandlePushCode(pkg *domain.PushCode) error {
	if pkg.Platform != domain.PlatformGithub {
		return nil
	}

	if !p.code.CheckRepoCreated(pkg.PkgName) {
		return fmt.Errorf("repo %s has not been created", pkg.PkgName)
	}

	repoUrl, err := p.code.Push(pkg)
	if err != nil {
		logrus.Errorf("pkgId %s push code err: %s", pkg.PkgId, err.Error())

		return err
	}

	e := domain.NewCodePushedEvent(pkg.PkgId, repoUrl)
	return p.producer.NotifyCodePushedResult(&e)
}
