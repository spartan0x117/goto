package cmd

import (
	"fmt"

	"github.com/pkg/browser"
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
		url := store.GetLinkForLabel(args[0])
		if url == "" {
			fmt.Printf("could not find label '%s'\n", args[0])
		}
		browser.OpenURL(url)
	},
}
