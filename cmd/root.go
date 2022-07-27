package cmd

import (
	"os"
	"os/signal"
	"syscall"

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

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
		signal.Notify(sigterm, syscall.SIGINT)
		<-sigterm
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
	// name of config file (without extension)
	viper.SetConfigName("testkube-watch")
	// adding $TKW_HOME directory as config search path
	if value, found := os.LookupEnv("TKW_HOME"); found {
		viper.AddConfigPath(value)
	}
	// adding $HOME directory as config search path
	if value, found := os.LookupEnv("HOME"); found {
		viper.AddConfigPath(value)
	}
	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file %s", viper.ConfigFileUsed())
	}
}
