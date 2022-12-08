package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pkg/browser"
)

type Goto struct {
	storage Storage
}

// TODO: Will need a way to initialize the directory ~/.config/goto/ and setup an initial config, possibly prompting the user for the github url to use?
func initializeGotoDirectory() {
	return
}

func LoadConfig() (*Config, error) {
	return loadConfig("~/.config/goto/config.yaml")
}

func NewGoto(c *Config) (Goto, error) {
	var s Storage
	switch c.Type {
	case "json":
		s = NewJsonStorage(c)
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

// TODO: Write a helper function to take any label and 'normalize' it, which should remove
// any '-' characters and return the resulting label

/*
Example commands (assuming that the executable is in $PATH and named 'goto'):

goto add agent-docs https://grafana.com/docs/agent/latest/        <---- Adds a link with the label 'agentdocs'

goto agent-docs                                                   <---- Tries to find a link with the label 'agentdocs', and opens it in the browser if it exists

goto sync                                                         <---- Pulls from github

goto find agent-docs                                              <---- Similar to basic goto, but only displays all matching links

goto remove agent-docs                                            <---- Removes 'agentdocs' link, if it exists

-- golinks.json

	{
		"agentdocs": "https://grafana.com/docs/agent/latest/",
		"bamboo": "https://grafana.bamboohr.com/"
	}
*/
func main() {
	args := os.Args[1:]
	numArgs := len(args)

	config, err := LoadConfig()
	if err != nil {
		fmt.Println("failed to load config")
	}
	go2, err := NewGoto(config)
	if err != nil {
		fmt.Println("failed to load goto from config")
		os.Exit(1)
	}

	switch args[0] {
	case "find":
		if numArgs == 0 {
			fmt.Println(go2.ListGotoLinks())
		} else if numArgs == 1 {
			fmt.Println(go2.SearchGotoLinks(args[1]))
		} else {
			fmt.Println("0 or 1 arguments for 'find'")
			os.Exit(1)
		}
	case "sync":
		if numArgs != 0 {
			fmt.Println("0 arguments for 'sync'")
			os.Exit(1)
		}
		fmt.Println(go2.Sync())
	case "remove":
		if numArgs != 1 {
			fmt.Println("1 argument for 'remove'")
			os.Exit(1)
		}
		fmt.Println(go2.RemoveGotoLink(args[1]))
	case "add":
		if numArgs != 2 {
			fmt.Println("2 arguments for 'add'")
			os.Exit(1)
		}
		go2.AddGotoLink(args[1], args[2])
	default:
		if numArgs != 1 {
			fmt.Println("goto expects 1 argument")
			os.Exit(1)
		}
		fmt.Println(go2.GotoLink(args[0]))
	}
}
