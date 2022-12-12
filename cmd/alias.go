package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

var aliasRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a local alias for an existing label",
	Long:  "Remove a local alias for an existing label. These configured aliases are stored in ~/.config/goto/aliases.json",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
