package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "htr",
	Short: "htr is a simple HTTP request CLI",
	Long:  "htr allows you to execute HTTP requests defined in a configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'htr run <config.yaml> [request_name]' to execute a request.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
