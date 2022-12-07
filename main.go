package main

import (
	"fmt"
	"os"
)

// TODO: Will need a way to initialize the directory ~/.config/goto/ and setup an initial config, possibly prompting the user for the github url to use as the initial?
func initializeGotoDirectory() {
	return
}

// TODO: Read the config from ~/.config/goto/config.yaml
func loadConfig() {
	return
}

// TODO: Write handler to sync (i.e. pull) for a particular github repo
func syncRepo(repoLabel string) error {
	return nil
}

// TODO: Write handler to add a new entry for a particular github repo
func addGotoLink(repoLabel string, linkLabel string, url string) error {
	return nil
}

// TODO: Write handler to remove a goto link for a particular repo
func removeGotoLink(repoLabel string, linkLabel string) error {
	return nil
}

// TODO: Write a handler to perform a search of goto links without actually visiting
func searchGotoLinks(repoLabel string, linkLabel string) string {
	return ""
}

// TODO: Write handler to goto a link, maybe use "github.com/pkg/browser"?
func gotoLink(repoLabel string, linkLabel string) error {
	return nil
}

func main() {
	args := os.Args[1:]

	fmt.Println(args)
}
