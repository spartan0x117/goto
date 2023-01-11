package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"
)

type JsonStorage struct {
	Path string `yaml:"path"`
}

func (jfs *JsonStorage) writeFile(links map[string]string) error {
	jsonString, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return errors.New("could not marshal link map")
	}
	return os.WriteFile(jfs.Path, jsonString, 0666)
}

func (jfs *JsonStorage) loadLinks() (map[string]string, error) {
	if !FileExists(jfs.Path) {
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
	l := m[label]
	if l != "" && !strings.HasPrefix(l, "https://") && !strings.HasPrefix(l, "http://") {
		l = "http://" + l
	}
	return l
}

func (jfs *JsonStorage) GetAllLabels() []string {
	m, err := jfs.loadLinks()
	if err != nil {
		return nil
	}

	links := make([]string, 0, len(m))
	for k := range m {
		links = append(links, k)
	}
	sort.Strings(links)
	return links
}

func (jfs *JsonStorage) AddLink(label string, url string, _ bool) error {
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
