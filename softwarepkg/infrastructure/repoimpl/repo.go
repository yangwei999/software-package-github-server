package repoimpl

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v36/github"
)

type iClient interface {
	CreateRepo(org string, r *github.Repository) error
}

func NewRepoImpl(cfg Config, cli iClient) *RepoImpl {
	return &RepoImpl{
		cfg: cfg,
		cli: cli,
	}
}

type RepoImpl struct {
	cfg Config
	cli iClient
}

func (impl *RepoImpl) CreateRepo(repo string) (string, error) {
	err := impl.cli.CreateRepo(
		impl.cfg.Org,
		&github.Repository{Name: &repo},
	)

	if err != nil && !strings.Contains(err.Error(), "name already exists") {
		return "", err
	}

	url := fmt.Sprintf("https://github.com/%s/%s", impl.cfg.Org, repo)

	return url, nil
}
