package errors

import "fmt"

// BaseError is the base error type, used to create other errors. Should not be used directly.
type BaseError struct {
	Message string `json:"reason"`
	Kind    string `json:"kind"`
	Param   string `json:"param"`
}

// Error implements the error interface.
func (e *BaseError) Error() string {
	return e.Message
}

// ErrAlreadyExists is returned when a resource already exists.
type ErrAlreadyExists struct {
	BaseError
}

// AlreadyExists returns a new ErrAlreadyExists.
func AlreadyExists(param string) error {
	return &ErrAlreadyExists{
		BaseError{
			Message: fmt.Sprintf("%s already exists", param),
			Kind:    "AlreadyExists",
		},
	}
}

// ErrInternalServerError is returned when an internal error occurs.
type ErrInternalServerError struct {
	BaseError
}

// InternalServerError returns a new ErrInternal.
func InternalServerError(reason string, args ...any) error {
	return &ErrInternalServerError{
		BaseError{
			Message: fmt.Sprintf(reason, args...),
			Kind:    "InternalServerError",
		},
	}
}

// ErrInvalid is returned when a parameter is invalid along with a reason.
type ErrInvalid struct {
	BaseError
}

// Invalid returns a new ErrInvalid.
func Invalid(param, reason string, args ...any) error {
	return &ErrInvalid{
		BaseError{
			Param:   param,
			Message: fmt.Sprintf(reason, args...),
			Kind:    "Invalid",
		},
	}
}

// ErrNoChanges is returned when no changes were made.
type ErrNoChanges struct {
	BaseError
}

// NoChanges returns a new ErrNoChanges.
func NoChanges() error {
	return &ErrNoChanges{
		BaseError{
			Message: "no changes were made",
			Kind:    "NoChanges",
		},
	}
}

// ErrNotFound is returned when a resource is not found.
type ErrNotFound struct {
	BaseError
}

// NotFound returns a new ErrNotFound.
func NotFound(param string) error {
	return &ErrNotFound{
		BaseError{
			Param:   param,
			Message: fmt.Sprintf("%s was not found", param),
			Kind:    "NotFound",
		},
	}
}

// ErrRequired is returned when a parameter is required.
type ErrRequired struct {
	BaseError
}

// Required returns a new ErrRequired.
func Required(param string) error {
	return &ErrRequired{
		BaseError{
			Param:   param,
			Message: fmt.Sprintf("%s is required", param),
			Kind:    "Required",
		},
	}
}

// ErrTimeout is returned when a timeout occurs.
type ErrTimeout struct {
	BaseError
}

// Timeout returns a new ErrTimeout.
func Timeout() error {
	return &ErrTimeout{
		BaseError{
			Message: "a timeout occurred, please try again",
			Kind:    "Timeout",
		},
	}
}

// ErrTooManyRequests is returned when a user has made too many requests.
type ErrTooManyRequests struct {
	BaseError
}

// TooManyRequests returns a new TooManyRequests.
func TooManyRequests() error {
	return &ErrTooManyRequests{
		BaseError{
			Message: "you have made too many requests, please try again later",
			Kind:    "TooManyRequests",
		},
	}
}

// ErrUnauthorized is returned when a user is not authorized.
type ErrUnauthorized struct {
	BaseError
}

// Unauthorized returns a new ErrUnauthorized.
func Unauthorized(reason string, args ...any) error {
	return &ErrUnauthorized{
		BaseError{
			Message: fmt.Sprintf(reason, args...),
			Kind:    "Unauthorized",
		},
	}
}

// ErrUnavailable is returned when a resource is unavailable.
type ErrUnavailable struct {
	BaseError
}

// Unavailable returns a new ErrUnavailable.
func Unavailable(param string) error {
	return &ErrUnavailable{
		BaseError{
			Message: fmt.Sprintf("%s is unavailable", param),
			Kind:    "Unavailable",
		},
	}
}
