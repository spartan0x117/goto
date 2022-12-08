package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Type       string `yaml:"type"`
	JsonConfig struct {
		Path string `yaml:"path"`
	} `yaml:"json_config,omitempty"`
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
