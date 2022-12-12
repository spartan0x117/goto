package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/browser"
	"github.com/spartan0x117/goto/pkg/alias"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open goto link in browser",
	Long:  "Open the goto link in the default browser",
	Run: func(cmd *cobra.Command, args []string) {
		label, path, pathExists := strings.Cut(args[0], "/")
		label = alias.GetLabelOrAlias(label)
		url := store.GetLinkForLabel(label)
		if url == "" {
			fmt.Printf("could not find label '%s'\n", args[0])
			os.Exit(1)
		}
		if pathExists {
			url = url + "/" + path
		}
		if err := browser.OpenURL(url); err != nil {
			fmt.Printf("error trying to open '%s' in browser\n", url)
			os.Exit(1)
		}
	},
}
