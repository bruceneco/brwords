package brwords

import "errors"

var (
	// ErrCantVisit means that the page could not be found or is not accessible.
	ErrCantVisit = errors.New("failed to visit page")
)
