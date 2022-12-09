package cmd

import (
	"fmt"
	"os"

	"github.com/spartan0x117/goto/pkg/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	store   storage.Storage
	rootCmd = &cobra.Command{
		Use:   "goto",
		Short: "goto is a tool to save and use urls for ease of use",
		Long:  "An OSS CLI version of go/ links that allows for either collaboratively building up a repository of useful links or just keeping your own",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/goto/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(fmt.Sprintf("%s/.config/goto", home))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not locate goto config file")
	}

	switch viper.GetString("type") {
	case "json":
		store = &storage.JsonStorage{
			Path: viper.GetString("json_config.path"),
		}
	}
}
