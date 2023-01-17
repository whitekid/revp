package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"revp/client"
	"revp/config"
)

var rootCmd = &cobra.Command{
	Use:   "revp local-address",
	Short: "revp",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		localAddr := args[0]

		client, err := client.New(localAddr, config.Client.ServerAddr())
		if err != nil {
			return err
		}

		ctx := cmd.Context()
		ctxClosed, cancel := context.WithCancel(context.Background())
		defer cancel()

		remoteAddr, err := client.Run(ctx, cancel)
		if err != nil {
			return err
		}

		fmt.Printf("forwarding %s -> %s\n", remoteAddr, localAddr)

		select {
		case <-ctx.Done(): // comand closed
		case <-ctxClosed.Done(): // proxy closed
		}

		return nil
	},
}

func init() {
	config.Init("revp", rootCmd.Flags())
	config.Client.Init("revp", rootCmd.Flags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
