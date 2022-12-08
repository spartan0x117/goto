package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"sort"
)

type JsonFileStorage struct {
	path string
}

func New(c Config) JsonFileStorage {
	return JsonFileStorage{
		path: c.JsonConfig.Path,
	}
}

func (jfs *JsonFileStorage) fileExists() bool {
	_, err := os.Stat(jfs.path)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}

func (jfs *JsonFileStorage) writeFile(links map[string]string) error {
	jsonString, err := json.Marshal(links)
	if err != nil {
		return errors.New("could not marshal link map")
	}
	return os.WriteFile(jfs.path, jsonString, 0666)
}

func (jfs *JsonFileStorage) loadLinks() (map[string]string, error) {
	if !jfs.fileExists() {
		return map[string]string{}, nil
	}
	m := map[string]string{}
	dat, err := os.ReadFile(jfs.path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(dat, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (jfs *JsonFileStorage) Sync() error {
	return nil
}

func (jfs *JsonFileStorage) GetLinkForLabel(label string) string {
	label = normalizeLabel(label)
	m, err := jfs.loadLinks()
	if err != nil {
		return ""
	}
	return m[label]
}

func (jfs *JsonFileStorage) GetAllLabels() []string {
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

func (jfs *JsonFileStorage) AddLink(label string, url string) error {
	label = normalizeLabel(label)
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

func (jfs *JsonFileStorage) RemoveLink(label string) error {
	label = normalizeLabel(label)
	m, err := jfs.loadLinks()
	if err != nil {
		return errors.New("could not unmarshal json file")
	}
	delete(m, label)
	return jfs.writeFile(m)
}
