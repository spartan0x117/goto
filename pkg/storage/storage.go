package storage

import (
	"errors"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

type Storage interface {
	// Pulls any updates from the underlying storage mechanism
	Sync() error

	// Gets the link for a particular label
	GetLinkForLabel(label string) string

	// Gets all label/link pairs
	GetAllLabels() []string

	// Adds a link for a label
	AddLink(label string, url string) error

	// Removes a link for a label
	RemoveLink(label string) error
}

func NormalizeLabel(label string) string {
	lower := strings.ToLower(label)
	re := regexp.MustCompile("[^a-z0-9]")
	return string(re.ReplaceAll([]byte(lower), []byte("")))
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}
