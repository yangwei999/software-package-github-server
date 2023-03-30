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

func NewPushCode(pkgName, importer, importerEmail, spec, rpm string) SoftwarePkg {
	return SoftwarePkg{
		PkgName: pkgName,
		Importer: Importer{
			Name:  importer,
			Email: importerEmail,
		},
		SourceCode: SourceCode{
			SpecURL:   spec,
			SrcRPMURL: rpm,
		},
	}
}
