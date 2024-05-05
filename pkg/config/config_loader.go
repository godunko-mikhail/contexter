package config

import (
	"github.com/godunko-mikhail/contexter/pkg/contexter"
	"gopkg.in/yaml.v2"
	"os"
)

func Load() (*Config, error) {
	text, err := os.ReadFile(contexter.GetConfigPath())
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(text, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Exists() bool {
	_, err := os.Stat(contexter.GetConfigPath())
	return err == nil
}
