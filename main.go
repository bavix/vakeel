package main

import (
	"context"

	"github.com/bavix/vakeel/cmd"
)

// main is the entry point of the application.
// It executes the command with an empty context.
// The context is used to pass the configuration to the command.
// The configuration is used to set the host and port of the server.
func main() {
	// Create an empty context.
	// The context is used to pass configuration settings to the command.
	ctx := context.Background()

	// Execute the command with the context.
	// The command is responsible for setting up and running the server.
	// It takes the context as a parameter and uses it to configure the server.
	cmd.Execute(ctx)
}
