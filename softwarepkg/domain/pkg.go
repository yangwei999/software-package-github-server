package domain

const (
	PkgStatusInitialized = "initialized"
	PkgStatusRepoCreated = "repo_created"
)

type Importer struct {
	Name  string
	Email string
}

type SourceCode struct {
	SpecURL   string
	SrcRPMURL string
}

type SoftwarePkgBasic struct {
	Id   string
	Name string
}

type SoftwarePkg struct {
	SoftwarePkgBasic

	Status   string
	Importer Importer
	SourceCode
	CIPRNum int
}

func (s *SoftwarePkg) SetPkgStatusRepoCreated() {
	s.Status = PkgStatusRepoCreated
}

func NewSoftwarePkg(b SoftwarePkgBasic, i Importer, s SourceCode, p int) SoftwarePkg {
	return SoftwarePkg{
		SoftwarePkgBasic: b,
		Importer:         i,
		SourceCode:       s,
		Status:           PkgStatusInitialized,
		CIPRNum:          p,
	}
}
