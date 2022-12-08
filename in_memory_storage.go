package main

import "sort"

type InMemoryStorage struct {
	linkLabelMap map[string]string
}

// noop, as the storage is all in memory
func (ims *InMemoryStorage) Sync() error {
	return nil
}

func (ims *InMemoryStorage) GetLinkForLabel(label string) string {
	label = normalizeLabel(label)
	return ims.linkLabelMap[label]
}

func (ims *InMemoryStorage) GetAllLabels() []string {
	labels := make([]string, 0, len(ims.linkLabelMap))
	for k := range ims.linkLabelMap {
		labels = append(labels, k)
	}
	sort.Strings(labels)
	return labels
}

func (ims *InMemoryStorage) AddLink(label string, url string) error {
	label = normalizeLabel(label)
	ims.linkLabelMap[label] = url
	return nil
}

func (ims *InMemoryStorage) RemoveLink(label string) error {
	label = normalizeLabel(label)
	delete(ims.linkLabelMap, label)
	return nil
}

func NewInMemoryStorage() *InMemoryStorage {
	ims := InMemoryStorage{
		linkLabelMap: map[string]string{},
	}
	return &ims
}
