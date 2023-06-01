package exception_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func TestItemSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ErrorSuite))
}

func (s *ErrorSuite) TestNewError() {
	type Sut struct {
		Sut  func() (*exception.Exception, error)
		Data *exception.Exception
	}

	makeSut := func() *Sut {
		data := fake.Exception()

		sut := func() (*exception.Exception, error) {
			return exception.New().
				SetID(data.ID).
				SetMessage(data.Message).
				SetStack(data.Stack).
				Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("create a new exception", func() {
		sut := makeSut()

		result, err := sut.Sut()

		s.NotNil(result)
		s.NoError(err)
	})
}
