package main

import (
	"os"

	"github.com/spf13/cobra"

	"revp/config"
	"revp/server"
)

var rootCmd = &cobra.Command{
	Use:   "revps",
	Short: "revps server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.New(config.Server.BindAddr()).Serve(cmd.Context())
	},
}

func init() {
	config.Init("revps", rootCmd.Flags())
	config.Server.Init("revps", rootCmd.Flags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
