package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spartan0x117/goto/pkg/storage"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Type       string              `yaml:"type"`
	JsonConfig storage.JsonStorage `yaml:"json_config,omitempty"`
	GitConfig  storage.GitStorage  `yaml:"git_config,omitempty"`
}

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not locate user home dir: %w", err)
	}
	return filepath.Join(home, ".config", "goto", "config.yaml"), nil
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
