//go:build linux || darwin

package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/bavix/vakeel/internal/build"
	"github.com/bavix/vakeel/internal/config"
	"github.com/bavix/vakeel/pkg/ctxid"
)

// init initializes the register command.
//
// It creates a new configuration object, sets up the register command, and adds it to the root command.
func init() {
	// Create a new configuration object.
	cfg := &config.Config{}

	// Create a new register command.
	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Register the agent with the server",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Create a new context with the ID value from the configuration.
			ctx := ctxid.WithID(cmd.Context(), cfg.ID)

			// Create a new builder with the configuration.
			builder := build.New(cfg)

			// Call the AgentRegisterApp method of the builder and pass the context of the command.
			// The AgentRegisterApp method registers the agent application with the server.
			return builder.AgentRegisterApp(builder.Logger(ctx))
		},
	}

	// Set the default value of the host flag to "127.0.0.1".
	registerCmd.Flags().
		StringVarP(&cfg.Host, "host", "H", "127.0.0.1", "Host for agent, i.e. the IP address of the Vakeel server.")

	// Set the default value of the port flag to 4643.
	registerCmd.Flags().
		IntVarP(&cfg.Port, "port", "p", 4643, "Port for agent, i.e. the port number of the Vakeel server.")

	// Set the default value of the id flag to a new UUID.
	// The flag is used to set the ID of the agent, i.e. the UUID of the Vakeel agent.
	// The UUID is generated using uuid.New() and converted to a string using uuid.String().
	registerCmd.Flags().
		StringVar(&cfg.ID, "id", uuid.New().String(), "ID of agent, i.e. the UUID of the Vakeel agent."+
			"If not provided, a new UUID will be generated.")

	// Add the register command to the root command.
	rootCmd.AddCommand(registerCmd)
}
