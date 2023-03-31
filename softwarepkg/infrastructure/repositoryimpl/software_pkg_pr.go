package repositoryimpl

import (
	"github.com/google/uuid"

	"github.com/opensourceways/software-package-github-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-github-server/softwarepkg/domain/repository"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/postgresql"
)

type softwarePkgPR struct {
	cli dbClient
}

func NewSoftwarePkgPR(cfg *Config) repository.SoftwarePkg {
	return softwarePkgPR{cli: postgresql.NewDBTable(cfg.Table.SoftwarePkgPR)}
}

func (s softwarePkgPR) Add(p *domain.SoftwarePkg) error {
	u, err := uuid.Parse(p.Id)
	if err != nil {
		return err
	}

	var do SoftwarePkgPRDO
	if err = s.toSoftwarePkgPRDO(p, u, &do); err != nil {
		return err
	}

	filter := SoftwarePkgPRDO{PkgId: u}

	return s.cli.Insert(&filter, &do)
}

func (s softwarePkgPR) Save(p *domain.SoftwarePkg) error {
	u, err := uuid.Parse(p.Id)
	if err != nil {
		return err
	}
	filter := SoftwarePkgPRDO{PkgId: u}

	var do SoftwarePkgPRDO
	if err = s.toSoftwarePkgPRDO(p, u, &do); err != nil {
		return err
	}

	return s.cli.UpdateRecord(&filter, &do)
}

func (s softwarePkgPR) FindAll() ([]domain.SoftwarePkg, error) {
	filter := SoftwarePkgPRDO{}

	var res []SoftwarePkgPRDO

	if err := s.cli.GetRecords(
		&filter,
		&res,
		postgresql.Pagination{},
		nil,
	); err != nil {
		return nil, err
	}

	var p = make([]domain.SoftwarePkg, len(res))

	for i := range res {
		v, err := res[i].toDomainPullRequest()
		if err != nil {
			return nil, err
		}

		p[i] = v
	}

	return p, nil
}

func (s softwarePkgPR) Remove(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	filter := SoftwarePkgPRDO{PkgId: u}

	return s.cli.DeleteRecord(&filter)
}
