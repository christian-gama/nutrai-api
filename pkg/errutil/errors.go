package errutil

import "fmt"

// ErrNotFound is returned when a resource is not found.
type ErrNotFound struct {
	Param string
}

// NotFound returns a new ErrNotFound.
//
// Returns a message like "could not find Param".
func NotFound(param string) error {
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

// Invalid returns a new ErrInvalid.
//
// Returns a message like "Param is invalid: Reason".
func Invalid(param, reason string, args ...any) error {
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

// Required returns a new ErrRequired.
//
// Returns a message like "Param is required".
func Required(param string) error {
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

// InternalServerError returns a new ErrInternal.
//
// Returns a message like "internal error: Reason".
func InternalServerError(reason string, args ...any) error {
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

// Unauthorized returns a new ErrUnauthorized.
//
// Returns a message like "unauthorized: Reason".
func Unauthorized(reason string, args ...any) error {
	return &ErrUnauthorized{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrUnauthorized) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Reason)
}

// ErrTooManyRequests is returned when a user has made too many requests.
type ErrTooManyRequests struct {
	Reason string
}

// TooManyRequests returns a new TooManyRequests.
//
// Returns a message like "too many requests: Reason".
func TooManyRequests(reason string, args ...any) error {
	return &ErrTooManyRequests{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrTooManyRequests) Error() string {
	return fmt.Sprintf("too many requests: %s", e.Reason)
}

// ErrRepository is returned when a repository error occurs.
type ErrRepository struct {
	Reason string
}

// Repository returns a new ErrRepository.
func Repository(reason string, args ...any) error {
	return &ErrRepository{Reason: fmt.Sprintf(reason, args...)}
}

// Error implements the error interface.
func (e *ErrRepository) Error() string {
	return fmt.Sprintf("repository error: %s", e.Reason)
}
