package codeimpl

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opensourceways/server-common-lib/utils"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
)

func NewCodeImpl(cfg Config) *CodeImpl {
	gitUrl := fmt.Sprintf(
		"https://%s:%s@github.com/%s/",
		cfg.Robot.Username,
		cfg.Robot.Token,
		cfg.Org,
	)

	repoUrl := fmt.Sprintf(
		"https://github.com/%s/",
		cfg.Org,
	)

	return &CodeImpl{
		gitUrl:  gitUrl,
		repoUrl: repoUrl,
		script:  cfg.ShellScript,
		watch:   cfg.Watch,
		ciRepo:  cfg.CIRepo,
		httpCli: utils.NewHttpClient(3),
	}
}

type CodeImpl struct {
	gitUrl  string
	repoUrl string
	script  string
	watch   Watch
	ciRepo  CIRepo
	httpCli utils.HttpClient
}

func (impl *CodeImpl) Push(pkg *domain.PushCode) (string, error) {
	repoUrl := fmt.Sprintf("%s%s.git", impl.gitUrl, pkg.PkgName)

	params := []string{
		impl.script,
		repoUrl,
		pkg.PkgName,
		pkg.Importer.Name,
		pkg.Importer.Email,
		impl.ciRepo.Link,
		impl.ciRepo.Repo,
	}

	out, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, out=%s, params=%v",
			err.Error(), string(out), params,
		)
	}

	return impl.repoUrl + pkg.PkgName, err
}

func (impl *CodeImpl) CheckRepoCreated(repo string) bool {
	repoUrl := fmt.Sprintf("%s%s.git", impl.gitUrl, repo)
	request, _ := http.NewRequest(http.MethodHead, repoUrl, nil)

	count := 0
	for {
		if code, _ := impl.httpCli.ForwardTo(request, nil); code == 0 {
			return true
		}

		count++
		if count > impl.watch.LoopTimes {
			return false
		}

		time.Sleep(impl.watch.IntervalDuration())
	}
}
