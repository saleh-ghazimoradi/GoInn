package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway"
	"github.com/saleh-ghazimoradi/GoInn/slg"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
		if err := gateway.Server(); err != nil {
			slg.Logger.Error(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
