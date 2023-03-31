package app

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/code"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/message"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repo"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repository"
)

type PkgService interface {
	HandleCreateRepo(*domain.SoftwarePkg) error
	HandlePushCode(*domain.SoftwarePkg) error
}

func NewPkgService(
	r repo.Repo, repo repository.SoftwarePkg,
	c code.Code, prod message.SoftwarePkgProducer,
) *pkgService {
	return &pkgService{
		repo:       r,
		code:       c,
		repository: repo,
		producer:   prod,
	}
}

type pkgService struct {
	repo       repo.Repo
	code       code.Code
	repository repository.SoftwarePkg
	producer   message.SoftwarePkgProducer
}

func (p *pkgService) HandleCreateRepo(pkg *domain.SoftwarePkg) error {
	url, err := p.repo.CreateRepo(pkg.Name)
	if err != nil {
		return err
	}

	e := domain.NewRepoCreatedEvent(pkg.Id, url)
	if err = p.producer.NotifyRepoCreatedResult(&e); err != nil {
		return err
	}

	pkg.SetPkgStatusRepoCreated()

	return p.repository.Save(pkg)
}

func (p *pkgService) HandlePushCode(pkg *domain.SoftwarePkg) error {
	repoUrl, err := p.code.Push(pkg)
	if err != nil {
		logrus.Errorf("pkgId %s push code err: %s", pkg.Id, err.Error())

		return err
	}

	e := domain.NewCodePushedEvent(pkg.Id, repoUrl)
	if err = p.producer.NotifyCodePushedResult(&e); err != nil {
		return err
	}

	return p.repository.Remove(pkg.Id)
}
