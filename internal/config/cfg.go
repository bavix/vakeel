package config

// Config holds the configuration for the vakeel agent.
type Config struct {
	// Host is the host address of the vakeel-way server.
	Host string
	// Port is the port number of the vakeel-way server.
	Port int
	// ID is the agent ID.
	ID string
}
