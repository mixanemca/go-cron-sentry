package app

import (
	"github.com/mixanemca/go-cron-sentry/models"
	"github.com/mixanemca/go-cron-sentry/runner"
	"github.com/mixanemca/go-cron-sentry/runner/local"
)

type app struct {
	dsn           string
	quiet         bool
	runnerBackend string

	task *models.Task

	runner runner.Client
}

// Option options for app
type Option func(c *app) error

// New creates a new App. Various client options can be used to configure
func New(opts ...Option) (App, error) {
	a := &app{
		task: &models.Task{},
	}
	var err error

	for _, opt := range opts {
		if err = opt(a); err != nil {
			return nil, err
		}
	}

	switch a.runnerBackend {
	case "local":
		a.runner, err = local.New(
			local.WithQuiet(a.quiet),
		)
	}
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Runner implements App interface
func (a *app) Runner() runner.Client {
	return a.runner
}
