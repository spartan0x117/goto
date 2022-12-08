package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pkg/browser"
)

type Goto struct {
	m LinkStorage
}

// TODO: Will need a way to initialize the directory ~/.config/goto/ and setup an initial config, possibly prompting the user for the github url to use?
func initializeGotoDirectory() {
	return
}

// TODO: Read the config from ~/.config/goto/config.yaml
func loadConfig() Config {
	return Config{
		Type: "todo",
	}
}

// TODO: Load the mapping from an actual file
func (c Config) LoadGoto() (Goto, error) {
	g := Goto{
		m: NewInMemoryStorage(),
	}
	g.AddGotoLink("gmail", "https://mail.google.com")
	g.AddGotoLink("grafana", "https://mail.google.com")
	return g, nil
}

// TODO: Write handler to sync (i.e. pull)
func (g *Goto) Sync() error {
	return nil
}

// TODO: Write handler to add a new entry
func (g *Goto) AddGotoLink(label string, url string) error {
	return nil
}

// TODO: Write handler to remove a goto link
func (g *Goto) RemoveGotoLink(label string) error {
	return nil
}

// TODO: Write a handler to perform a search of goto links without actually visiting
func (g *Goto) SearchGotoLinks(label string) string {
	return g.m.GetLinkForLabel(label)
}

// TODO: Write a handler to list all goto links
func (g *Goto) ListGotoLinks() []string {
	return g.m.GetAllLabels()
}

// TODO: Write handler to goto a link, maybe use "github.com/pkg/browser"?
func (g *Goto) GotoLink(label string) error {
	url := g.m.GetLinkForLabel(label)
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

	config := loadConfig()
	go2, err := config.LoadGoto()
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
