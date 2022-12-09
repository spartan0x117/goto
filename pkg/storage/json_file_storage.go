package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"sort"
)

type JsonStorage struct {
	Path string `yaml:"path"`
}

func (jfs *JsonStorage) fileExists() bool {
	_, err := os.Stat(jfs.Path)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}

func (jfs *JsonStorage) writeFile(links map[string]string) error {
	jsonString, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return errors.New("could not marshal link map")
	}
	return os.WriteFile(jfs.Path, jsonString, 0666)
}

func (jfs *JsonStorage) loadLinks() (map[string]string, error) {
	if !jfs.fileExists() {
		return map[string]string{}, nil
	}
	m := map[string]string{}
	dat, err := os.ReadFile(jfs.Path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(dat, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (jfs *JsonStorage) Sync() error {
	return nil
}

func (jfs *JsonStorage) GetLinkForLabel(label string) string {
	label = NormalizeLabel(label)
	m, err := jfs.loadLinks()
	if err != nil {
		return ""
	}
	return m[label]
}

func (jfs *JsonStorage) GetAllLabels() []string {
	m, err := jfs.loadLinks()
	if err != nil {
		return nil
	}

	labels := make([]string, 0, len(m))
	for k := range m {
		labels = append(labels, k)
	}
	sort.Strings(labels)
	return labels
}

func (jfs *JsonStorage) AddLink(label string, url string) error {
	label = NormalizeLabel(label)
	m, err := jfs.loadLinks()
	if err != nil {
		return errors.New("could not unmarshal json file")
	}
	m[label] = url
	err = jfs.writeFile(m)
	if err != nil {
		return errors.New("could not marshal file after adding label/link pair")
	}
	return nil
}

func (jfs *JsonStorage) RemoveLink(label string) error {
	label = NormalizeLabel(label)
	m, err := jfs.loadLinks()
	if err != nil {
		return errors.New("could not unmarshal json file")
	}
	delete(m, label)
	return jfs.writeFile(m)
}
