package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryStorageGet(t *testing.T) {
	ims := InMemoryStorage{
		linkLabelMap: map[string]string{
			"gmail":   "https://mail.google.com",
			"example": "https://example.com",
		},
	}

	assert.Equal(t, "https://mail.google.com", ims.GetLinkForLabel("gmail"))
	assert.Equal(t, "https://example.com", ims.GetLinkForLabel("example"))
}

func TestInMemoryStorageGetAllLabels(t *testing.T) {
	ims := InMemoryStorage{
		linkLabelMap: map[string]string{
			"gmail":   "https://mail.google.com",
			"example": "https://example.com",
		},
	}

	expected := []string{"example", "gmail"}
	assert.Equal(t, expected, ims.GetAllLabels())
}

func TestInMemoryStorageAddLink(t *testing.T) {
	ims := InMemoryStorage{linkLabelMap: map[string]string{}}

	assert.NoError(t, ims.AddLink("grafana", "https://grafana.com"))
	assert.Contains(t, ims.linkLabelMap, "grafana")
	assert.Equal(t, "https://grafana.com", ims.linkLabelMap["grafana"])
}

func TestInMemoryStorageRemoveLink(t *testing.T) {
	ims := InMemoryStorage{
		linkLabelMap: map[string]string{
			"gmail":   "https://mail.google.com",
			"example": "https://example.com",
		},
	}

	assert.NoError(t, ims.RemoveLink("gmail"))
	assert.NotContains(t, ims.linkLabelMap, "gmail")
	assert.Contains(t, ims.linkLabelMap, "example")
}
