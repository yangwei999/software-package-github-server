package repoimpl

type Config struct {
	Org string `json:"org"`
}

func (cfg *Config) SetDefault() {

	cfg.Org = "orgfsdfsd"

}
