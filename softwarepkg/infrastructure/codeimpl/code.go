package codeimpl

import (
	"fmt"
	"strconv"

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
		ciRepo:  cfg.CIRepo,
	}
}

type CodeImpl struct {
	gitUrl  string
	repoUrl string
	script  string
	ciRepo  CIRepo
}

func (impl *CodeImpl) Push(pkg *domain.PushCode) (string, error) {
	repoUrl := fmt.Sprintf("%s%s.git", impl.gitUrl, pkg.PkgName)

	params := []string{
		impl.script,
		repoUrl,
		pkg.PkgName,
		pkg.Importer,
		pkg.ImporterEmail,
		impl.ciRepo.Link,
		impl.ciRepo.Repo,
		strconv.Itoa(pkg.CIPRNum),
	}

	out, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, out=%s, params=%v",
			err.Error(), string(out), params[:len(params)-1],
		)
	}

	return impl.repoUrl + pkg.PkgName, err
}
