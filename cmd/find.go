package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds goto links in the store according to the input regex",
	Long:  "Finds goto links in the store according to the input regex. If no regex is supplied it will return all goto links in the store",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Println(store.GetAllLabels())
		case 1:
			fmt.Println(store.GetLinkForLabel(args[0]))
		}
	},
}
