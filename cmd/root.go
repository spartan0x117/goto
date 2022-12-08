package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goto",
	Short: "goto is a tool to save and use urls for ease of use",
	Long:  "A OSS CLI version of go/ links that allows for either collaboratively building up a repository of useful links or just keeping your own",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd)
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
