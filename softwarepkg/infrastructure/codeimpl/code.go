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

	return &CodeImpl{
		gitUrl: gitUrl,
		script: cfg.ShellScript,
	}
}

type CodeImpl struct {
	gitUrl string
	script string
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
	}

	_, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, params=%v",
			err.Error(), params[:len(params)-1],
		)
	}

	return repoUrl, err
}
