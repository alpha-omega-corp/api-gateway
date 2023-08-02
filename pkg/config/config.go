package config

import (
	"embed"
	"gopkg.in/yaml.v3"
	"io/fs"
	"sync"
)

var (
	//go:embed envs
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

type Config struct {
	HOST    string `yaml:"host"`
	AUTH    string `yaml:"auth"`
	PRODUCT string `yaml:"product"`
}

func LoadConfig() (*Config, error) {
	return readConfig(makeFS(), "dev")
}

func readConfig(fileSys fs.FS, env string) (*Config, error) {
	b, err := fs.ReadFile(fileSys, env+".yaml")
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func makeFS() fs.FS {
	unwrapFSOnce.Do(func() {
		fileSys, err := fs.Sub(embedFS, "envs")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fileSys
	})
	return unwrappedFS
}
