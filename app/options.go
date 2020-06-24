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
