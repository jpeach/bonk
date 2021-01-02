package cli

import (
	"fmt"
	"os"
)

const (
	// EX_OK is a successful program exit code.
	EX_OK = 0 //nolint(golint)

	// EX_FAIL is a generic failure program exit code.
	EX_FAIL = 1 //nolint(golint)
)

// Execute runs f and exits based on the error result.
func Execute(progname string, f func() error) {
	if err := f(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", progname, err.Error())
		os.Exit(EX_FAIL)
	}

	os.Exit(EX_OK)
}
