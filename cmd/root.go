package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		configDir := fmt.Sprintf("%s/.config/goto", home)
		cobra.CheckErr(err)

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		err = os.MkdirAll(configDir, os.ModePerm)
		// don't fail if directory already exists
		if err != nil && !os.IsExist(err) {
			panic(fmt.Errorf("could not create %v: %w", configDir, err))
		}
		configPath := filepath.Join(configDir, "config.yaml")
		if _, err := os.Stat(configPath); err != nil {
			// default container dir for links.json is configDir
			jsonStoragePath := filepath.Join(configDir, "links.json")
			viper.SetDefault("type", "json")
			viper.SetDefault("json_config", map[string]string{"path": jsonStoragePath})
			err = viper.WriteConfigAs(configPath)
			if err != nil {
				panic(fmt.Errorf("%v: %w", configPath, err))
			}
		}
		aliasesPath := filepath.Join(configDir, "aliases.json")
		if _, err := os.Stat(aliasesPath); err != nil {
			err := os.WriteFile(aliasesPath, []byte("{}\n"), 0600)
			if err != nil {
				panic(fmt.Errorf("%v: %w", aliasesPath, err))
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not locate goto config file")
	}

	switch viper.GetString("type") {
	case "json":
		store = &storage.JsonStorage{
			Path: viper.GetString("json_config.path"),
		}
	case "git":
		store = &storage.GitStorage{
			LocalPath: viper.GetString("git_config.local_path"),
			AutoSync:  viper.GetBool("git_config.auto_sync"),
		}
	}
}
