package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	jfs.AddLink("grafana", "https://grafana.com", false)
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

	jfs.AddLink("grafana", "https://grafana.com", false)
	assert.NoError(t, jfs.RemoveLink("grafana"))
	assert.Equal(t, "", jfs.GetLinkForLabel("grafana"))
}
