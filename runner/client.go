package runner

type client struct {
	quite bool
}

// Option represtns option for Runner client
type Option func(c *client) error

// New returns Runner
func New(opts ...Option) (Client, error) {
	c := &client{}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
