package app

// WithDSN sets an app's DSN
func WithDSN(dsn string) Option {
	return func(a *app) error {
		a.dsn = dsn
		return nil
	}
}

// WithQuiet suppress all command output
func WithQuiet(quiet bool) Option {
	return func(a *app) error {
		a.quiet = quiet
		return nil
	}
}

// WithTask make a cron task
func WithTask(command string, args ...string) Option {
	return func(a *app) error {
		a.task.Command = command
		a.task.Args = args
		return nil
	}
}

// WithRunnerBackend sets an app's Runner backend
func WithRunnerBackend(runner string) Option {
	return func(a *app) error {
		a.runnerBackend = runner
		return nil
	}
}
