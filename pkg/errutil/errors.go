package errutil

import "fmt"

// ErrNotFound is returned when a resource is not found.
type ErrNotFound struct {
	Param string
}

// NewErrNotFound returns a new ErrNotFound.
//
// Returns a message like "could not find Param".
func NewErrNotFound(param string) error {
	return &ErrNotFound{Param: param}
}

// Error implements the error interface.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("could not find %s", e.Param)
}

// ErrInvalid is returned when a parameter is invalid along with a reason.
type ErrInvalid struct {
	Param  string
	Reason string
}

// NewErrInvalid returns a new ErrInvalid.
//
// Returns a message like "Param is invalid: Reason".
func NewErrInvalid(param, reason string, args ...any) error {
	return &ErrInvalid{Param: param, Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrInvalid) Error() string {
	return fmt.Sprintf("%s is invalid: %s", e.Param, e.Reason)
}

// ErrRequired is returned when a parameter is required.
type ErrRequired struct {
	Param string
}

// NewErrRequired returns a new ErrRequired.
//
// Returns a message like "Param is required".
func NewErrRequired(param string) error {
	return &ErrRequired{Param: param}
}

// Error implements the error interface.
func (e *ErrRequired) Error() string {
	return fmt.Sprintf("%s is required", e.Param)
}

// ErrInternal is returned when an internal error occurs.
type ErrInternal struct {
	Reason string
}

// NewErrInternal returns a new ErrInternal.
//
// Returns a message like "internal error: Reason".
func NewErrInternal(reason string, args ...any) error {
	return &ErrInternal{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrInternal) Error() string {
	return fmt.Sprintf("internal error: %s", e.Reason)
}

// ErrUnauthorized is returned when a user is not authorized.
type ErrUnauthorized struct {
	Reason string
}

// NewErrUnauthorized returns a new ErrUnauthorized.
//
// Returns a message like "unauthorized: Reason".
func NewErrUnauthorized(reason string, args ...any) error {
	return &ErrUnauthorized{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrUnauthorized) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Reason)
}

// TooManyRequests is returned when a user has made too many requests.
type TooManyRequests struct {
	Reason string
}

// NewErrTooManyRequests returns a new TooManyRequests.
//
// Returns a message like "too many requests: Reason".
func NewErrTooManyRequests(reason string, args ...any) error {
	return &TooManyRequests{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *TooManyRequests) Error() string {
	return fmt.Sprintf("too many requests: %s", e.Reason)
}
