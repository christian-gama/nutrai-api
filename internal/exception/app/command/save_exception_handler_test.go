package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/app/command"
	"github.com/christian-gama/nutrai-api/testutils/jsonutil"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type SaveExceptionHandlerSuite struct {
	suite.Suite
}

func TestSaveExceptionHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SaveExceptionHandlerSuite))
}

func (s *SaveExceptionHandlerSuite) TestHandle() {
	type Mock struct {
		Publisher *messageMock.Publisher
	}

	type Sut struct {
		Sut   command.SaveExceptionHandler
		Ctx   context.Context
		Input *command.SaveExceptionInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Publisher: messageMock.NewPublisher(s.T()),
		}

		input := fake.SaveExceptionInput()

		sut := command.NewSaveExceptionHandler(mock.Publisher)

		return &Sut{Sut: sut, Ctx: context.Background(), Input: input, Mock: mock}
	}

	s.Run("publishes a new error", func() {
		sut := makeSut()

		sut.Mock.Publisher.
			On("Handle", jsonutil.MustMarshal(s.T(), sut.Input)).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})
}
