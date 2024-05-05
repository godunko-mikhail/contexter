package config

import (
	"errors"
)

type Config struct {
	Contexts []*Context `yaml:"contexts,omitempty"`
}

func (c *Config) Validate() error {

	return nil
}

func (c *Config) GetContext(name string) (*Context, error) {
	for _, ctx := range c.Contexts {
		if ctx.Name == name {
			return ctx, nil
		}
	}
	return nil, errors.New("context not found")
}
