package suite

import (
	"github.com/christian-gama/nutrai-api/testutils/suite/asserts"
	"github.com/stretchr/testify/suite"
)

// Suite is the base suite for all test suites. It provides helper methods for
// testing.
type Suite struct {
	suite.Suite
}

// Skip skips the test.
func (s *Suite) Skip(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skip()
	})
}

// Todo marks the test as TODO, skipping it and printing a message.
func (s *Suite) Todo(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skipf("TODO: %s", name)
	})
}

func (s *Suite) Run(name string, f func()) bool {
	return s.Suite.Run(name, func() {
		f()
	})
}

// ErrorAsAlreadyExists checks if the error is an ErrAlreadyExists.
func (s *Suite) ErrorAsAlreadyExists(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsAlreadyExists(s.T(), err, msgAndArgs...)
}

// ErrorAsInternalServerError checks if the error is an ErrInternalServerError.
func (s *Suite) ErrorAsInternalServerError(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInternalServerError(s.T(), err, msgAndArgs...)
}

// ErrorAsInvalid checks if the error is an ErrInvalid.
func (s *Suite) ErrorAsInvalid(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInvalid(s.T(), err, msgAndArgs...)
}

// ErrorAsNoChanges checks if the error is an ErrNoChanges.
func (s *Suite) ErrorAsNoChanges(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsNoChanges(s.T(), err, msgAndArgs...)
}

// ErrorAsNotFound checks if the error is an ErrNotFound.
func (s *Suite) ErrorAsNotFound(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsNotFound(s.T(), err, msgAndArgs...)
}

// ErrorAsRequired checks if the error is an ErrRequired.
func (s *Suite) ErrorAsRequired(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsRequired(s.T(), err, msgAndArgs...)
}

// ErrorAsTimeout checks if the error is an ErrTimeout.
func (s *Suite) ErrorAsTimeout(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsTimeout(s.T(), err, msgAndArgs...)
}

// ErrorAsTooManyRequests checks if the error is an ErrTooManyRequests.
func (s *Suite) ErrorAsTooManyRequests(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsTooManyRequests(s.T(), err, msgAndArgs...)
}

// ErrorAsUnauthorized checks if the error is an ErrUnauthorized.
func (s *Suite) ErrorAsUnauthorized(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsUnauthorized(s.T(), err, msgAndArgs...)
}

// ErrorAsUnavailable checks if the error is an ErrUnavailable.
func (s *Suite) ErrorAsUnavailable(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsUnavailable(s.T(), err, msgAndArgs...)
}

// HasNotChanged checks if the given objects are equal.
func (s *Suite) HasNotChanged(
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	return asserts.HasNotChanged(s.T(), oldObj, newObj, msgAndArgs...)
}

// HasChanged checks if the given object has changed.
func (s *Suite) HasChanged(
	oldObj, newObj any,
	msgAndArgs ...any,
) bool {
	return asserts.HasChanged(s.T(), oldObj, newObj, msgAndArgs...)
}
