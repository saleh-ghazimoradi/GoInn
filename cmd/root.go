package cmd

import (
	"github.com/saleh-ghazimoradi/GoInn/config"
	"github.com/saleh-ghazimoradi/GoInn/slg"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoInn",
	Short: "GoInn, a Hotel Reservation App",
}

func Execute() {
	err := os.Setenv("TZ", time.UTC.String())
	if err != nil {
		panic(err)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := config.LoadConfig()
	if err != nil {
		slg.Logger.Error("Failed to load config", "error", err)
		return
	}
}
