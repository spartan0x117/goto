package cmd

import (
	"fmt"

	"github.com/spartan0x117/goto/pkg/alias"
	"github.com/spf13/cobra"
)

func init() {
	aliasCmd.AddCommand(aliasAddCmd)
	aliasCmd.AddCommand(aliasRemoveCmd)
	rootCmd.AddCommand(aliasCmd)
}

var aliasCmd = &cobra.Command{
	Use:       "alias",
	Short:     "Add or remove a local alias for an existing label",
	ValidArgs: []string{"add", "remove"},
}

var aliasAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a local alias for an existing label",
	Long:  "Add a local alias for an existing label. These configured aliases are stored in ~/.config/goto/aliases.json",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := alias.AddAlias(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	},
}

var aliasRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a local alias for an existing label",
	Long:  "Remove a local alias for an existing label. These configured aliases are stored in ~/.config/goto/aliases.json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := alias.RemoveAlias(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}
