package main

import (
	"errors"

	"github.com/pkg/browser"
	"github.com/spartan0x117/goto/cmd"
	"github.com/spartan0x117/goto/pkg/config"
	"github.com/spartan0x117/goto/pkg/storage"
)

type Goto struct {
	storage storage.Storage
}

func NewGoto(c *config.Config) (Goto, error) {
	var s storage.Storage
	switch c.Type {
	case "json":
		s = config.NewJsonStorage(c)
	}

	g := Goto{
		storage: s,
	}
	return g, nil
}

func (g *Goto) Sync() error {
	return g.storage.Sync()
}

func (g *Goto) AddGotoLink(label string, url string) error {
	return g.storage.AddLink(label, url)
}

func (g *Goto) RemoveGotoLink(label string) error {
	return g.storage.RemoveLink(label)
}

func (g *Goto) SearchGotoLinks(label string) string {
	return g.storage.GetLinkForLabel(label)
}

func (g *Goto) ListGotoLinks() []string {
	return g.storage.GetAllLabels()
}

func (g *Goto) GotoLink(label string) error {
	url := g.storage.GetLinkForLabel(label)
	if url == "" {
		return errors.New("could not find label")
	}
	return browser.OpenURL(url)
}

func main() {
	cmd.Execute()
}
