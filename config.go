package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Type       string      `yaml:"type"`
	JsonConfig JsonStorage `yaml:"json_config,omitempty"`
}

func loadConfig(path string) (*Config, error) {
	c := &Config{}
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(fileContents, c)
	return c, nil
}