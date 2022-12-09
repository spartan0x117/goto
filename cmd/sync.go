package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Pulls goto links from configured source",
	Long:  "Pulls goto links from configured source. This is effectively a no-op for the local json file store",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Make this a prettier message
		fmt.Println(store.Sync())
	},
}
