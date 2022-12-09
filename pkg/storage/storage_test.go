package storage

import (
	"os"
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

func TestJsonFileStorageGet(t *testing.T) {
	jfs := JsonStorage{
		Path: "./testdata/sample_links.json",
	}
	actual := jfs.GetLinkForLabel("grafana")
	assert.Equal(t, "https://grafana.com", actual)
}

func TestJsonFileStorageGetAllLabels(t *testing.T) {
	jfs := JsonStorage{
		Path: "./testdata/sample_links.json",
	}

	expected := []string{"gmail", "grafana"}
	assert.Equal(t, expected, jfs.GetAllLabels())
}

func TestJsonFileStorageAddLink(t *testing.T) {
	jfs := JsonStorage{
		Path: "./testdata/tmp.json",
	}
	t.Cleanup(func() {
		os.Remove("./testdata/tmp.json")
	})

	jfs.AddLink("grafana", "https://grafana.com")
	actual := jfs.GetLinkForLabel("grafana")
	assert.Equal(t, "https://grafana.com", actual)
}

func TestJsonFileStorageRemoveLink(t *testing.T) {
	jfs := JsonStorage{
		Path: "./testdata/tmp.json",
	}
	t.Cleanup(func() {
		os.Remove("./testdata/tmp.json")
	})

	jfs.AddLink("grafana", "https://grafana.com")
	assert.NoError(t, jfs.RemoveLink("grafana"))
	assert.Equal(t, "", jfs.GetLinkForLabel("grafana"))
}
