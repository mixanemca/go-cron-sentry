package runner

// WithQuiet suppress all command output
func WithQuiet(quiet bool) Option {
	return func(c *client) error {
		c.quite = quiet
		return nil
	}
}
