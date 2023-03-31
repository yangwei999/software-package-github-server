package config

import (
	"github.com/opensourceways/server-common-lib/utils"

	"github.com/opensourceways/software-package-github-server/message-server"
	"github.com/opensourceways/software-package-github-server/mq"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/codeimpl"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/postgresql"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/repoimpl"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/repositoryimpl"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/watchingimpl"
)

type PostgresqlConfig struct {
	DB postgresql.Config `json:"db" required:"true"`

	repositoryimpl.Config
}

type Config struct {
	MQ            mq.Config            `json:"mq"`
	Postgresql    PostgresqlConfig     `json:"postgresql"`
	MessageServer messageserver.Config `json:"message_server"`
	Repo          repoimpl.Config      `json:"repo"`
	Code          codeimpl.Config      `json:"code"`
	Watch         watchingimpl.Config  `json:"watch"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := new(Config)
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return nil, err
	}

	cfg.SetDefault()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.MQ,
		&cfg.Postgresql,
		&cfg.MessageServer,
		&cfg.Repo,
		&cfg.Code,
		&cfg.Watch,
	}
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

func (cfg *Config) SetDefault() {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) Validate() error {
	if _, err := utils.BuildRequestBody(cfg, ""); err != nil {
		return err
	}

	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
