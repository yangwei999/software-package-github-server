package codeimpl

import (
	"errors"
	"fmt"
	"os/exec"

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

func (impl *CodeImpl) Push(pkg *domain.SoftwarePkg) error {
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

	err := RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, params=%v",
			err.Error(), params[:len(params)-1],
		)
	}

	return err
}

func RunCmd(args ...string) error {
	n := len(args)
	if n == 0 {
		return nil
	}

	cmd := args[0]

	if n > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	c := exec.Command(cmd, args...)
	out, err := c.CombinedOutput()
	if err != nil {
		return errors.New(string(out) + err.Error())
	}

	return nil
}
