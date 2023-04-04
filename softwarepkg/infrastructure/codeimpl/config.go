package codeimpl

type Config struct {
	ShellScript string      `json:"shell_script"`
	Org         string      `json:"org"`
	Robot       RobotConfig `json:"robot"`
}

func (c *Config) SetDefault() {
	if c.ShellScript == "" {
		c.ShellScript = "/opt/app/code.sh"
	}

	if c.Org == "" {
		c.Org = "src-openeuler"
	}
}

type RobotConfig struct {
	Username string `json:"username" required:"true"`
	Token    string `json:"token"    required:"true"`
}
