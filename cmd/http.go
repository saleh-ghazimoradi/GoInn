package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "runs the http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
