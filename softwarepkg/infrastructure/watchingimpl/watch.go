package watchingimpl

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/softwarepkg/app"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repository"
)

func NewWatchingImpl(
	c Config, r repository.SoftwarePkg, p app.PkgService,
) *WatchingImpl {
	return &WatchingImpl{
		cfg:        c,
		repo:       r,
		pkgService: p,
	}
}

type WatchingImpl struct {
	cfg        Config
	repo       repository.SoftwarePkg
	pkgService app.PkgService
}

func (impl *WatchingImpl) Start(ctx context.Context, stop chan struct{}) {
	interval := impl.cfg.IntervalDuration()

	checkStop := func() bool {
		select {
		case <-ctx.Done():
			return true
		default:
			return false
		}
	}

	for {
		pkgs, err := impl.repo.FindAll()
		if err != nil {
			logrus.Errorf("find all storage pkg failed, err: %s", err.Error())
		}

		for _, pkg := range pkgs {
			impl.handle(&pkg)
		}

		if checkStop() {
			close(stop)

			return
		}
		time.Sleep(interval)
	}
}

func (impl *WatchingImpl) handle(pkg *domain.SoftwarePkg) {
	switch pkg.Status {
	case domain.PkgStatusInitialized:
		if err := impl.pkgService.HandleCreateRepo(pkg); err != nil {
			logrus.Errorf("handle create repo err: %s", err.Error())
		}
	case domain.PkgStatusRepoCreated:
		if err := impl.pkgService.HandlePushCode(pkg); err != nil {
			logrus.Errorf("handle push code err: %s", err.Error())
		}
	}
}
