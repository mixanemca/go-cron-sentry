package app

import "github.com/mixanemca/go-cron-sentry/runner"

// App is the root-level interface for interacting with the OS exec and Senty.
// You can instantiate an implementation of this interface using the "New" function.
type App interface {
	// Runner returns a specialized API for interacting with OS for executing commands
	Runner() runner.Client
	// Reporter returns a specialized API for interacting with Sentry
	// Reporter() reporter.Client
}
