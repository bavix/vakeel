package app

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	apiv1 "github.com/bavix/apis/pkg/bavix/api/v1"
	"github.com/bavix/apis/pkg/uuidconv"
	"github.com/bavix/vakeel-way/pkg/api/vakeel_way"
	"github.com/bavix/vakeel/pkg/ctxid"
)

// duration is the duration between update requests sent by the agent.
// It is set to 15 seconds.
//
// Update requests are sent continuously to the server with this duration.
// The agent will send an update request to the server every 15 seconds.
const duration = 15 * time.Second

// Agent sends update requests to the given state service client.
// It continuously sends update requests until the context is cancelled.
// Returns an error if sending the update request fails.
//
// Parameters:
// - ctx: The context.Context to use for the gRPC call.
// - stateServiceClient: The client for the state service.
//
// Returns:
// - error: An error if sending the update request fails.
func Agent(
	ctx context.Context,
	stateServiceClient vakeel_way.StateServiceClient,
) error {
	// Loop until the context is cancelled.
	for {
		select {
		case <-ctx.Done():
			// If the context is cancelled, return the error from the context.
			return ctx.Err()
		default:
			// Create a client stream to send updates to the server.
			// This method establishes a connection with the server and returns a client stream.
			// If the connection fails, an error is returned.
			updateClient, err := stateServiceClient.Update(ctx)
			if err != nil {
				// Log the error and sleep for a duration before continuing.
				logError(ctx, err, "failed to create client stream")
				time.Sleep(duration)

				continue
			}

			// Send an update request to the server.
			// This function sends an update request to the server using the client stream.
			// If sending the update request fails, an error is returned.
			if err := stream(ctx, updateClient); err != nil {
				// Log the error and continue.
				logError(ctx, err, "failed to send update request")
			}

			// Close the update stream to free resources.
			// This method closes the client stream and waits for the response from the server.
			// If the response is not received, an error is returned.
			if _, err := updateClient.CloseAndRecv(); err != nil {
				// Log the error and continue.
				logError(ctx, err, "failed to close update stream")
			}
		}
	}
}

// logError logs the error with the given message.
//
// It takes a context, an error, and a message as parameters.
// The function logs the error with the given message using the zerolog logger.
// The logger is obtained from the context and the error is logged with the message.
//
// Parameters:
// - ctx: The context.Context used for logging.
// - err: The error to log.
// - msg: The message to log along with the error.
func logError(ctx context.Context, err error, msg string) {
	// Get the logger from the context.
	logger := zerolog.Ctx(ctx)

	// Log the error with the message.
	logger.Error().Err(err).Msg(msg)
}

// stream sends an update request to the server at regular intervals.
//
// It takes a context and a client for the update service as parameters.
// The function sends an update request to the server with the ID extracted from the context.
// It also logs a message indicating that an update request is being sent.
// The function returns an error if sending the update request fails.
func stream(
	ctx context.Context,
	client vakeel_way.StateService_UpdateClient,
) error {
	// Get the high and low parts of the UUID from the context.
	// The UUID is extracted from the context using the ctxid.ID function.
	high, low := uuidconv.UUID2DoubleInt(ctxid.ID(ctx))

	// Send an initial update request to the server with the given UUID.
	// The sendUpdateRequest function logs a message indicating that an update request is being sent
	// and returns an error if sending the update request fails.
	if err := sendUpdateRequest(ctx, client, high, low); err != nil {
		return err
	}

	// Create a ticker to send update requests at regular intervals.
	// The duration between update requests is set to 15 seconds.
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	// Loop until the context is cancelled.
	for {
		select {
		// If the context is cancelled, return nil.
		case <-ctx.Done():
			return nil

		// If the ticker fires, send an update request to the server with the given UUID.
		case <-ticker.C:
			// The sendUpdateRequest function logs a message indicating that an update request is being sent
			// and returns an error if sending the update request fails.
			if err := sendUpdateRequest(ctx, client, high, low); err != nil {
				return err
			}
		}
	}
}

// sendUpdateRequest sends an update request to the server with the given UUID.
// It takes a context, a client for the server's update service, and the high and low parts of the UUID.
// The function logs a message indicating that an update request is being sent
// and returns an error if sending the update request fails.
func sendUpdateRequest(
	ctx context.Context,
	client vakeel_way.StateService_UpdateClient,
	high int64,
	low int64,
) error {
	// Create an update request with the UUID extracted from the context.
	updateRequest := &vakeel_way.UpdateRequest{
		Ids: []*apiv1.UUID{
			{High: high, Low: low},
		},
	}

	// Log a message indicating that an update request is being sent.
	// The message includes the UUID that is being sent.
	zerolog.Ctx(ctx).Info().Msgf("sending update request: %s", uuidconv.DoubleInt2UUID(high, low))

	// Send the update request to the server.
	// The function returns an error if sending the update request fails.
	return client.Send(updateRequest)
}
