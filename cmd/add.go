package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new goto link",
	Long:  "Add a new goto link with the specified label and optional description",
	Run: func(cmd *cobra.Command, args []string) {
		store.AddLink(args[0], args[1])
	},
}
