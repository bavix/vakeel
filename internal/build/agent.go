package build

import (
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bavix/vakeel-way/pkg/api/vakeel_way"
	"github.com/bavix/vakeel/internal/app"
	"github.com/bavix/vakeel/internal/infra/templater"
	"github.com/bavix/vakeel/pkg/ctxid"
)

// AgentApp creates a gRPC client and connects to the server's update service.
// It returns an error if the connection or the update service call fails.
//
// ctx: The context.Context to use for the gRPC call.
// Returns: An error if the connection or update service call fails.
func (b *Builder) AgentApp(ctx context.Context) error {
	// Create a gRPC client insecure connection to the server.
	conn, err := grpc.NewClient(
		net.JoinHostPort(b.config.Host, strconv.Itoa(b.config.Port)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	// Close the connection when the function returns.
	defer conn.Close()

	serviceClient := vakeel_way.NewStateServiceClient(conn)

	return app.Agent(ctx, serviceClient)
}

// AgentRegisterApp is a method of the Builder struct.
//
// It registers the agent application with the server.
// It returns an error if the registration fails.
//
// ctx: The context.Context to use for the gRPC call.
//
// Returns:
// An error if the registration fails.
func (b *Builder) AgentRegisterApp(ctx context.Context) error {
	// Create a new templater.New instance with the context ID, host, and port from the config.
	// The templater.New instance generates the stub agent template.
	generate, err := templater.New(ctxid.ID(ctx), b.config.Host, b.config.Port)
	if err != nil {
		return err
	}

	// Call the AgentRegister function of the app package.
	// It creates a gRPC client insecure connection to the server.
	// It registers the agent application with the server.
	// It returns an error if the registration fails.
	//
	// ctx: The context.Context to use for the gRPC call.
	// generate: The templater.New instance that generates the stub agent template.
	//
	// Returns:
	// An error if the registration fails.
	return app.AgentRegister(ctx, generate)
}
