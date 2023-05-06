package errutil_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func TestErrorSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ErrorSuite))
}

func (s *ErrorSuite) TestError() {
	s.Run("Error.Len returns the number of errors correctly", func() {
		var result *errutil.Error
		result = errutil.Append(result, errors.New("error 1"))
		result = errutil.Append(result, errors.New("error 2"))

		s.Equal(2, result.Len())
	})

	s.Run("Error.HasErrors returns true if there are errors", func() {
		var result *errutil.Error
		result = errutil.Append(result, errors.New("error 1"))
		result = errutil.Append(result, errors.New("error 2"))

		s.Equal(true, result.HasErrors())
	})

	s.Run("Error.HasErrors returns false if there are no errors", func() {
		var result *errutil.Error

		s.Equal(false, result.HasErrors())
	})

	s.Run("Error.Error returns the error message correctly", func() {
		var result *errutil.Error
		result = errutil.Append(result, errors.New("error 1"))
		result = errutil.Append(result, errors.New("error 2"))

		s.Equal(
			fmt.Sprintf("occurred %d errors:\n\t- error 1\n\t- error 2", result.Len()),
			result.Error(),
		)
	})

	s.Run("errors.Unwrap returns the first error", func() {
		err := errors.New("error 1")

		var e *errutil.Error
		e = errutil.Append(e, err)
		e = errutil.Append(e, errors.New("error 2"))

		s.Equal(err.Error(), errors.Unwrap(e).Error())
	})

	s.Run("errors.Is returns true if the error is in the list", func() {
		err := errors.New("error 1")
		err2 := errors.New("error 2")

		var e *errutil.Error
		e = errutil.Append(e, err)

		s.Equal(true, errors.Is(e, err))
		s.Equal(false, errors.Is(e, err2))
	})
}
