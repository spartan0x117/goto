package cmd

import (
	"fmt"

	"github.com/spartan0x117/goto/pkg/alias"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds and displays goto links without opening in the browser",
	Long:  "Finds and displays the link for the goto label specified. If no label is supplied it will display all goto labels in the store. If there is a large number, recommended to pipe the output to less or grep",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			labels := store.GetAllLabels()
			for _, l := range labels {
				fmt.Println(l)
			}
		case 1:
			fmt.Println(store.GetLinkForLabel(alias.GetLabelOrAlias(args[0])))
		}
	},
}
