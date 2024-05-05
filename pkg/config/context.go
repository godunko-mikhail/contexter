package config

type Context struct {
	Name       string            `yaml:"name,omitempty"`
	Executable string            `yaml:"executable,omitempty"`
	Envs       map[string]string `yaml:"envs,omitempty"`
	Commands   []string          `yaml:"commands,omitempty"`
}
