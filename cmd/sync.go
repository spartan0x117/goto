package cmd

import (
	"fmt"
	"os"

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
		if err := store.Sync(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
