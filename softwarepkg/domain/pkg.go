package domain

type PushCode struct {
	Importer      string `json:"importer"`
	ImporterEmail string `json:"importer_email"`
	PkgId         string `json:"pkg_id"`
	PkgName       string `json:"pkg_name"`
	Platform      string `json:"platform"`
	CIPRNum       int    `json:"ci_pr_num"`
}
