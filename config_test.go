package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig("./testdata/sample_config.yaml")
	assert.NoError(t, err)

	assert.Equal(t, "json", c.Type)
	assert.Equal(t, "./testdata/sample_links.json", c.JsonConfig.Path)
}
