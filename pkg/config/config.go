package config

import (
	"os"

	"github.com/spartan0x117/goto/pkg/storage"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Type       string              `yaml:"type"`
	JsonConfig storage.JsonStorage `yaml:"json_config,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(fileContents, c)
	return c, nil
}

func NewJsonStorage(c *Config) *storage.JsonStorage {
	return &storage.JsonStorage{
		Path: c.JsonConfig.Path,
	}
}
