package cmd

import (
	"strings"

	"github.com/68696c6c/goat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var RootCommand = &cobra.Command{
	Use:   "horns",
	Short: "Launch the specified process.",
	Long:  "Entry point for the app.  Used to run all app processes.",
}

func init() {
	cobra.OnInitialize(initGoat)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("configFile", "./config.yml")
	RootCommand.PersistentFlags().StringVar(&configFile, "config", "./config.yml", "config file (default is ./config.yml)")
	viper.SetDefault("author", "Aaron Hill <68696c6c@gmail.com>")
	viper.SetDefault("license", "MIT")
}

func initGoat() {
	goat.ReadConfig(false)
	goat.Init()
}
