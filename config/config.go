package config

import (
	kafka "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/server-common-lib/utils"

	"github.com/opensourceways/software-package-github-server/message-server"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/codeimpl"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/useradapterimpl"
)

type Config struct {
	Kafka         kafka.Config           `json:"kafka"`
	MessageServer messageserver.Config   `json:"message_server"`
	Code          codeimpl.Config        `json:"code"`
	OmApi         useradapterimpl.Config `json:"om_api"`
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
		&cfg.Kafka,
		&cfg.MessageServer,
		&cfg.Code,
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
