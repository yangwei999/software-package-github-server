package codeimpl

import "time"

type Config struct {
	ShellScript string      `json:"shell_script"`
	Org         string      `json:"org"`
	Watch       Watch       `json:"watch"`
	Robot       RobotConfig `json:"robot"   required:"true"`
	CIRepo      CIRepo      `json:"ci_repo" required:"true"`
}

func (c *Config) SetDefault() {
	if c.Watch.Interval == 0 {
		c.Watch.Interval = 10
	}

	if c.Watch.LoopTimes == 0 {
		c.Watch.LoopTimes = 10
	}

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

type CIRepo struct {
	Repo string `json:"repo" required:"true"`
	Link string `json:"link" required:"true"`
}

type Watch struct {
	Interval  int `json:"interval"`
	LoopTimes int `json:"loop_times"`
}

func (c *Watch) IntervalDuration() time.Duration {
	return time.Second * time.Duration(c.Interval)
}
