package config

import (
	"errors"
)

type Config struct {
	Contexts []*Context `yaml:"contexts"`
}

func (c *Config) validate() error {

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

func (c *Config) AddContext(ctx *Context) error {
	if ctx == nil {
		return errors.New("context is nil")
	}

	if err := ctx.Validate(); err != nil {
		return err
	}

	c.Contexts = append(c.Contexts, ctx)
	return nil
}
