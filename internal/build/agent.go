package build

import (
	"context"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/bavix/vakeel-way/pkg/api/vakeel_way"
	"github.com/bavix/vakeel/internal/app"
	"github.com/bavix/vakeel/internal/infra/templater"
	"github.com/bavix/vakeel/pkg/ctxid"
)

// Keep-alive parameters for the gRPC client.
// The client will send a ping to the server every 10 seconds if there is no activity.
// The client will consider the connection dead if it doesn't receive a pong within 1 second.
// The client will allow the connection to be established without a stream.
const (
	keepAliveTime       = 10 * time.Second // The time between pings.
	keepAliveTimeout    = 1 * time.Second  // The time to wait for a pong.
	allowWithoutStreams = true             // Allow the connection to be established without a stream.
)

// AgentApp creates a gRPC client and connects to the server's update service.
// It returns an error if the connection or the update service call fails.
//
// ctx: The context.Context to use for the gRPC call.
// Returns: An error if the connection or update service call fails.
func (b *Builder) AgentApp(ctx context.Context) error {
	// Create a gRPC client insecure connection to the server.
	// The connection is established using the host and port from the configuration.
	// The connection is configured with keep-alive parameters to send pings to the server
	// every 10 seconds if there is no activity and to consider the connection dead if
	// a ping ack is not received within 1 second.
	conn, err := grpc.NewClient(
		net.JoinHostPort(b.config.Host, strconv.Itoa(b.config.Port)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                keepAliveTime,
			Timeout:             keepAliveTimeout,
			PermitWithoutStream: allowWithoutStreams,
		}),
	)
	if err != nil {
		return err
	}

	// Close the connection when the function returns.
	defer conn.Close()

	// Create a client for the vakeel_way.StateService.
	serviceClient := vakeel_way.NewStateServiceClient(conn)

	// Call the app.Agent function to start the agent.
	// The agent sends update requests to the server using the client stream.
	// The function returns an error if sending the update request fails.
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
