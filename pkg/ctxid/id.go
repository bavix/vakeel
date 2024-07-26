package ctxid

import (
	"context"

	"github.com/google/uuid"
)

// idKey is the key type used to store the ID value in the context.
type idKey struct{}

// WithID returns a new context with the provided ID value.
//
// ctx: The parent context.
// id: The ID value to attach to the context.
// Returns: A new context with the ID value attached.
func WithID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, &idKey{}, uuid.MustParse(id))
}

// ID retrieves the ID value from the provided context.
//
// ctx: The context to retrieve the ID from.
// Returns: The ID value from the context, or an empty UUID if the context doesn't have an ID.
func ID(ctx context.Context) uuid.UUID {
	// Get the value associated with the idKey from the context.
	// The value is either a uuid.UUID or nil.
	value := ctx.Value(&idKey{})

	// If the value is not a uuid.UUID, return an empty UUID.
	// This is the default value for the ID if it is not set in the context.
	vid, ok := value.(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	// Return the value as a uuid.UUID.
	// This is the value that is associated with the idKey in the context.
	return vid
}
