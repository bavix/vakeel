package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

//nolint:exhaustruct
var rootCmd = &cobra.Command{
	Use:   "vakeel",
	Short: "Agent for vakeel-way",
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
