package main

import (
	"os"

	"github.com/spartan0x117/goto/cmd"
)

func main() {
	validCommands := map[string]bool{
		"add":    true,
		"alias":  true,
		"find":   true,
		"init":   true,
		"open":   true,
		"remove": true,
		"server": true,
		"sync":   true,
	}

	// Only need to check the case where one argument has been supplied
	// beyond the filename. This inserts the 'open' command into the
	// supplied arguments as a hack to get 'goto mywebsite' to not error
	// in cobra
	if len(os.Args) == 2 {
		_, validCommand := validCommands[os.Args[1]]
		if !validCommand {
			os.Args = []string{
				os.Args[0],
				"open",
				os.Args[1],
			}
		}
	}

	cmd.Execute()
}
