package domain

type Importer struct {
	Name  string
	Email string
}

type SourceCode struct {
	SpecURL   string
	SrcRPMURL string
}

type SoftwarePkg struct {
	PkgName string
	Importer
	SourceCode
}

func NewSoftwarePkg(pkgName string, i Importer, s SourceCode) SoftwarePkg {
	return SoftwarePkg{
		PkgName:    pkgName,
		Importer:   i,
		SourceCode: s,
	}
}
