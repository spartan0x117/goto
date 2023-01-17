package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config files",
	Long:  "Initializes goto config files, if they don't exist",
	Run: func(cmd *cobra.Command, args []string) {
		// do nothing, since the initialization happens in root.go
	},
}
