package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var configLoader struct {
	path string
}

type Loader interface {
	Load() (*Config, error)
}

const (
	DefaultConfigFileName = ".contexter.conf"
)

func Load() (*Config, error) {
	text, err := os.ReadFile(configLoader.path)
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

func InitLoader() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	var sb strings.Builder
	sb.WriteString(home)
	sb.WriteString(string(os.PathSeparator))
	sb.WriteString(DefaultConfigFileName)

	configLoader.path = sb.String()
}
