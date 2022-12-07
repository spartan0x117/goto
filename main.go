package main

import (
	"fmt"
	"os"
)

type Config struct {
	gitRepoPath string
}

type GotoLinkMap map[string]string

// TODO: Will need a way to initialize the directory ~/.config/goto/ and setup an initial config, possibly prompting the user for the github url to use?
func initializeGotoDirectory() {
	return
}

// TODO: Read the config from ~/.config/goto/config.yaml
func loadConfig() {
	return
}

// TODO: Load the mapping from an actual file
func loadGotoLinkMap(repoPath string) (GotoLinkMap, error) {
	return nil, nil
}

// TODO: Write handler to sync (i.e. pull)
func sync() error {
	return nil
}

// TODO: Write handler to add a new entry
func addGotoLink(label string, url string) error {
	return nil
}

// TODO: Write handler to remove a goto link
func removeGotoLink(label string) error {
	return nil
}

// TODO: Write a handler to perform a search of goto links without actually visiting
func searchGotoLinks(label string) string {
	return ""
}

// TODO: Write a handler to list all goto links
func listGotoLinks() []string {
	return []string{}
}

// TODO: Write handler to goto a link, maybe use "github.com/pkg/browser"?
func gotoLink(label string) error {
	return nil
}

// TODO: Write a helper function to take any label and 'normalize' it, which should remove
// any '-' characters and return the resulting label
func normalizeLabel(label string) string {
	return label
}

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
	numArgs := len(args) - 1

	switch args[0] {
	case "find":
		if numArgs == 0 {
			fmt.Println(listGotoLinks())
		} else if numArgs == 1 {
			fmt.Println(searchGotoLinks(args[1]))
		} else {
			fmt.Println("0 or 1 arguments for 'find'")
			os.Exit(1)
		}
	case "sync":
		if numArgs != 0 {
			fmt.Println("0 arguments for 'sync'")
			os.Exit(1)
		}
		fmt.Println(sync())
	case "remove":
		if numArgs != 1 {
			fmt.Println("1 argument for 'remove'")
			os.Exit(1)
		}
		fmt.Println(removeGotoLink(args[1]))
	case "add":
		if numArgs != 2 {
			fmt.Println("2 arguments for 'add'")
			os.Exit(1)
		}
		addGotoLink(args[1], args[2])
	default:
		if numArgs != 1 {
			fmt.Println("goto expects 1 argument")
			os.Exit(1)
		}
		gotoLink(args[1])
	}
}
