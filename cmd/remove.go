package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a goto link",
	Long:  "Removes a goto link from the configured store and updates it",
	Run: func(cmd *cobra.Command, args []string) {
		store.RemoveLink(args[0])
	},
}
