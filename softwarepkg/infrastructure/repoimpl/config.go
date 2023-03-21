package repoimpl

type Config struct {
	Org string `json:"org"`
}

func (cfg *Config) SetDefault() {
	if cfg.Org == "" {
		cfg.Org = "src-openeuler"
	}
}
