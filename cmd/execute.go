package cmd

import (
	"fmt"
	"os"

	"github.com/joaovds/htr/internal/config"
	"github.com/joaovds/htr/internal/request"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <path/to/config.yaml> [request_name]",
	Short: "Executes a request from the configuration file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noStyle, _ := cmd.Flags().GetBool("no-style")
		filename := args[0]

		cfg, err := config.LoadConfig(filename)
		if err != nil {
			fmt.Println("Error loading config:", err)
			os.Exit(1)
		}

		if len(args) < 2 {
			fmt.Println("Available requests:")
			for name := range cfg.Requests {
				fmt.Println("-", name)
			}
			os.Exit(0)
		}

		reqName := args[1]
		reqConfig, exists := cfg.Requests[reqName]
		if !exists {
			fmt.Println("Request not found:", reqName)
			os.Exit(1)
		}

		req := request.New(cfg.BaseURL, reqConfig, noStyle)

		if err := req.Run(); err != nil {
			fmt.Println("Request failed:", err)
			os.Exit(1)
		}
	},
}

func init() {
	runCmd.Flags().BoolP("no-style", "n", false, "Exit without stylization")
	rootCmd.AddCommand(runCmd)
}
