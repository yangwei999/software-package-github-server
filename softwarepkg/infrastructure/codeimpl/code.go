package codeimpl

import (
	"fmt"

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
		token:   cfg.Robot.Token,
	}
}

type CodeImpl struct {
	gitUrl  string
	repoUrl string
	script  string
	token   string
}

func (impl *CodeImpl) Push(pkg *domain.SoftwarePkg) (string, error) {
	repoUrl := fmt.Sprintf("%s%s.git", impl.gitUrl, pkg.Name)

	params := []string{
		impl.script,
		repoUrl,
		pkg.Name,
		pkg.Importer.Name,
		pkg.Importer.Email,
		pkg.SpecURL,
		pkg.SrcRPMURL,
		impl.token,
	}

	out, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, out=%s, params=%v",
			err.Error(), string(out), params[:len(params)-1],
		)
	}

	return impl.repoUrl + pkg.Name, err
}
