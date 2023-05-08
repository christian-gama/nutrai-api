package errutil

import (
	"fmt"
)

// Error is a collection of errors.
type Error struct {
	Errors []error
}

// Error implements error.
func (e *Error) Error() string {
	if !e.HasErrors() {
		return ""
	}

	output := fmt.Sprintf("occurred %d errors:\n", len(e.Errors))
	for i, err := range e.Errors {
		output += fmt.Sprintf("\t- %s", err.Error())

		if i < len(e.Errors)-1 {
			output += "\n"
		}
	}

	return output
}

// HasErrors returns true if the error contains errors.
func (e *Error) HasErrors() bool {
	if e == nil || len(e.Errors) == 0 {
		return false
	}

	return true
}

// Len returns the number of errors.
func (e *Error) Len() int {
	if !e.HasErrors() {
		return 0
	}

	return len(e.Errors)
}

// Unwrap returns the first error in the chain.
func (e *Error) Unwrap() error {
	if !e.HasErrors() {
		return nil
	}

	if len(e.Errors) == 1 {
		return e.Errors[0]
	}

	errs := make([]error, len(e.Errors))
	copy(errs, e.Errors)
	return chain(errs)
}
