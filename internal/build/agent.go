package build

import (
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bavix/vakeel-way/pkg/api/vakeel_way"
	"github.com/bavix/vakeel/internal/app"
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
