package code

import "github.com/opensourceways/software-package-github-server/softwarepkg/domain"

type Code interface {
	Push(*domain.SoftwarePkg) (string, error)
}
