package app

import "github.com/mixanemca/go-cron-sentry/runner"

type app struct {
	dsn   string
	quiet bool

	runner runner.Client
}

// Option options for app
type Option func(c *app) error

// New creates a new App. Various client options can be used to configure
func New(opts ...Option) (App, error) {
	a := &app{}
	var err error

	for _, opt := range opts {
		if err = opt(a); err != nil {
			return nil, err
		}
	}

	a.runner, err = runner.New(
		runner.WithQuiet(a.quiet),
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Runner implements App interface
func (a *app) Runner() runner.Client {
	return a.runner
}
