package errutil

import "errors"

// chain is a slice of errors that implements the error interface.
// It is useful for using errors.Is and errors.As to check for a specific error
// in a chain of errors.
type chain []error

// Error returns the first error in the chain.
func (e chain) Error() string {
	return e[0].Error()
}

// Unwrap returns the next error in the chain.
func (e chain) Unwrap() error {
	if len(e) == 1 {
		return nil
	}

	return e[1:]
}

// As finds the first error in the chain that matches target, and if so, sets.
func (e chain) As(target any) bool {
	return errors.As(e[0], target)
}

// Is reports whether any error in the chain matches target.
func (e chain) Is(target error) bool {
	return errors.Is(e[0], target)
}
