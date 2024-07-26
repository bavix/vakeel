package build

import "github.com/bavix/vakeel/internal/config"

// Builder is a structure that provides methods to build the agent and the client.
type Builder struct {
	// config is the configuration for the agent and the client.
	config *config.Config
}

// New creates a new Builder instance with the given configuration.
// It returns a pointer to the Builder instance.
func New(config *config.Config) *Builder {
	return &Builder{
		config: config,
	}
}
