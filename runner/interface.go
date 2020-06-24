package runner

// Client defines the interface for Runner client
type Client interface {
	// Run execute command with args
	Run(command string, args ...string) (string, error)
}
