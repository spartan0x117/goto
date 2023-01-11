package cmd

import (
	"github.com/spf13/cobra"

	"github.com/spartan0x117/goto/pkg/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server running locally",
	Long:  "Starts a server on port 80 running locally.",
	Run: func(cmd *cobra.Command, args []string) {
		e := server.NewServer(store)
		e.Logger.Fatal(e.Start(":80"))
	},
}
