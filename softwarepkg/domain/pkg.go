package domain

type PushCode struct {
	Importer          string `json:"importer"`
	ImporterEmail     string `json:"importer_email"`
	PkgId             string `json:"pkg_id"`
	PkgName           string `json:"pkg_name"`
	PkgDesc           string `json:"pkg_desc"`
	ImportingPkgSig   string `json:"sig"`
	ReasonToImportPkg string `json:"reason_to_import"`
	Platform          string `json:"platform"`
	CIPRNum           int    `json:"ci_pr_num"`
}
