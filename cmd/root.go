package cmd

import (
	"fmt"

	"github.com/lreimer/testkube-watch-controller/config"
	"github.com/lreimer/testkube-watch-controller/pkg/controller"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "testkube-watch",
	Short: "Run Testkube test on Kubernetes events",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}
		controller.Start(config)
		fmt.Scanln()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Disable Help subcommand
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

func initConfig() {
	viper.SetConfigFile(config.DefaultConfigFileName)
	viper.SetConfigName("testkube-watch") // name of config file (without extension)
	viper.AddConfigPath("$TKW_HOME")      // adding $TKW_HOME directory as first search path
	viper.AddConfigPath("$HOME")          // adding $HOME directory as second search path
	viper.AutomaticEnv()                  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file %s", viper.ConfigFileUsed())
	}
}
