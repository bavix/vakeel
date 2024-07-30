package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vakeel",
	Short: "Agent for vakeel-way",
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalln(err)
	}
}
