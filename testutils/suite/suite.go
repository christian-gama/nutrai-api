package suite

import (
	"github.com/christian-gama/nutrai-api/testutils/suite/asserts"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) Skip(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skip()
	})
}

func (s *Suite) Todo(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skipf("TODO: %s", name)
	})
}

// ErrorAsNotFound checks if the error is an ErrNotFound.
func (s *Suite) ErrorAsNotFound(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsNotFound(s.T(), err, msgAndArgs...)
}

// ErrorAsInvalid checks if the error is an ErrInvalid.
func (s *Suite) ErrorAsInvalid(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInvalid(s.T(), err, msgAndArgs...)
}

// ErrorAsRequired checks if the error is an ErrRequired.
func (s *Suite) ErrorAsRequired(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsRequired(s.T(), err, msgAndArgs...)
}

// ErrorAsInternal checks if the error is an ErrInternal.
func (s *Suite) ErrorAsInternal(err error, msgAndArgs ...any) bool {
	return asserts.ErrorAsInternal(s.T(), err, msgAndArgs...)
}
