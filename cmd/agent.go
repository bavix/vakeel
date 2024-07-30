package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/bavix/vakeel/internal/build"
	"github.com/bavix/vakeel/internal/config"
	"github.com/bavix/vakeel/pkg/ctxid"
)

// init registers the agent command to the root command.
//
// The agent command is responsible for running the Vakeel agent. It establishes
// a connection to the Vakeel server and starts sending update requests.
func init() {
	// Create a new configuration object.
	cfg := &config.Config{}

	// Create a new agent command.
	agentCmd := &cobra.Command{
		Use:   "agent",
		Short: "Run the Vakeel agent",
		// RunE is the function that will be executed when the agent command is called.
		// It creates a new builder with the configuration and calls the AgentApp method of the builder.
		// The AgentApp method establishes a connection to the Vakeel server and starts sending update requests.
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Create a new context with the ID value from the configuration.
			ctx := ctxid.WithID(cmd.Context(), cfg.ID)

			// Create a new builder with the configuration.
			builder := build.New(cfg)

			// Call the AgentApp method of the builder and pass the context of the command.
			// The AgentApp method returns an error if the connection or the update service call fails.
			return builder.AgentApp(builder.Logger(ctx))
		},
	}

	// Set the default value of the host flag to "127.0.0.1".
	agentCmd.Flags().
		StringVarP(&cfg.Host, "host", "H", "127.0.0.1", "Host for agent, i.e. the IP address of the Vakeel server.")

	// Set the default value of the port flag to 4643.
	agentCmd.Flags().
		IntVarP(&cfg.Port, "port", "p", 4643, "Port for agent, i.e. the port number of the Vakeel server.")

	// Set the default value of the id flag to uuid.Nil.String().
	agentCmd.Flags().
		StringVar(&cfg.ID, "id", uuid.Nil.String(), "ID of agent, i.e. the UUID of the Vakeel agent.")

	// Add the agent command to the root command.
	rootCmd.AddCommand(agentCmd)
}
