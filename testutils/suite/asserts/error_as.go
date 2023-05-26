package asserts

import (
	"testing"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/stretchr/testify/assert"
)

// ErrorAsAlreadyExists asserts that the error is an ErrAlreadyExists.
func ErrorAsAlreadyExists(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrAlreadyExists
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsInternalServerError asserts that the error is an ErrInternalServerError.
func ErrorAsInternalServerError(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrInternalServerError
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsInvalid asserts that the error is an ErrInvalid.
func ErrorAsInvalid(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrInvalid
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsNoChanges asserts that the error is an ErrNoChanges.
func ErrorAsNoChanges(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrNoChanges
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsNotFound asserts that the error is an ErrNotFound.
func ErrorAsNotFound(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrNotFound
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsRequired asserts that the error is an ErrRequired.
func ErrorAsRequired(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrRequired
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsTimeout asserts that the error is an ErrTimeout.
func ErrorAsTimeout(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrTimeout
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsTooManyRequests asserts that the error is an ErrTooManyRequests.
func ErrorAsTooManyRequests(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrTooManyRequests
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsUnauthorized asserts that the error is an ErrUnauthorized.
func ErrorAsUnauthorized(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrUnauthorized
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}

// ErrorAsUnavailable asserts that the error is an ErrUnavailable.
func ErrorAsUnavailable(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errors.ErrUnavailable
	return assert.ErrorAs(t, err, &e, msgAndArgs...)
}
