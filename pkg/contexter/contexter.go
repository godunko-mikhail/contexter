package contexter

import (
	"os"
	"path"
	"path/filepath"
)

var (
	configPath string
)

func GetConfigPath() string {
	if configPath == "" {
		panic("bug: init required!")
	}
	return configPath
}

func Init(configPathOption string) {
	switch {
	case configPathOption != "":
		p, err := filepath.Abs(configPathOption)
		if err != nil {
			panic(err)
		}
		configPath = p
	default:
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configPath = path.Join(home, ".contexter.conf")
	}
}
