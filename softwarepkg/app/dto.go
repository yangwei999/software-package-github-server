package app

type CmdToCreateRepo struct {
	PkgId    string `json:"pkg_id"`
	PkgName  string `json:"pkg_name"`
	Platform string `json:"platform"`
}
