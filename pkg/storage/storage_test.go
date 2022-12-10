package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizedLabels(t *testing.T) {
	uppercase := "UpperCaseLabel"
	assert.Equal(t, "uppercaselabel", NormalizeLabel(uppercase))

	dashes := "-grafana-agent-"
	assert.Equal(t, "grafanaagent", NormalizeLabel(dashes))

	numbers := "spartan0x117"
	assert.Equal(t, numbers, NormalizeLabel(numbers))
}

func TestFileExists(t *testing.T) {
	exists := "./storage_test.go"
	assert.True(t, FileExists(exists))

	doesNotExist := "./does_not_exist"
	assert.False(t, FileExists(doesNotExist))
}
