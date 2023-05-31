package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/app/command"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type CatchExceptionHandlerSuite struct {
	suite.Suite
}

func TestCatchExceptionHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CatchExceptionHandlerSuite))
}

func (s *CatchExceptionHandlerSuite) TestHandle() {
	type Mock struct {
		Publisher *messageMock.Publisher[exception.Exception]
	}

	type Sut struct {
		Sut   command.CatchExceptionHandler
		Ctx   context.Context
		Input *command.CatchExceptionInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Publisher: messageMock.NewPublisher[exception.Exception](s.T()),
		}

		input := fake.CatchExceptionInput()

		sut := command.NewCatchExceptionHandler(mock.Publisher)

		return &Sut{Sut: sut, Ctx: context.Background(), Input: input, Mock: mock}
	}

	s.Run("publishes a new error", func() {
		sut := makeSut()

		sut.Mock.Publisher.
			On("Handle", mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
		sut.Mock.Publisher.AssertCalled(
			s.T(),
			"Handle",
			mock.MatchedBy(func(e exception.Exception) bool {
				return e.Stack == sut.Input.Stack && e.Message == sut.Input.Message
			}),
		)
	})
}
