package asserts

import (
	"testing"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/stretchr/testify/assert"
)

// ErrorAsNotFound asserts that the error is as ErrNotFound.
func ErrorAsNotFound(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errutil.ErrNotFound
	return assert.ErrorAs(t, err, &e, msgAndArgs)
}

// ErrorAsInvalid asserts that the error is as ErrInvalid.
func ErrorAsInvalid(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errutil.ErrInvalid
	return assert.ErrorAs(t, err, &e, msgAndArgs)
}

// ErrorAsRequired asserts that the error is as ErrRequired.
func ErrorAsRequired(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errutil.ErrRequired
	return assert.ErrorAs(t, err, &e, msgAndArgs)
}

// ErrorAsInternal asserts that the error is as ErrInternal.
func ErrorAsInternal(t *testing.T, err error, msgAndArgs ...any) bool {
	var e *errutil.ErrInternal
	return assert.ErrorAs(t, err, &e, msgAndArgs)
}
