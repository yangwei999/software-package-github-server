package repository

import "github.com/opensourceways/software-package-github-server/softwarepkg/domain"

type SoftwarePkg interface {
	Add(pkg *domain.SoftwarePkg) error
	Save(*domain.SoftwarePkg) error
	FindAll() ([]domain.SoftwarePkg, error)
	Remove(string) error
}
