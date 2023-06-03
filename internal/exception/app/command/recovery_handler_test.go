package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/app/command"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type RecoveryHandlerSuite struct {
	suite.Suite
}

func TestRecoveryHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RecoveryHandlerSuite))
}

func (s *RecoveryHandlerSuite) TestHandle() {
	type Mock struct {
		Publisher *messageMock.Publisher[command.RecoveryInput]
	}

	type Sut struct {
		Sut   command.RecoveryHandler
		Ctx   context.Context
		Input *command.RecoveryInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Publisher: messageMock.NewPublisher[command.RecoveryInput](s.T()),
		}

		input := fake.RecoveryInput()

		sut := command.NewRecoveryHandler(mock.Publisher)

		return &Sut{Sut: sut, Ctx: context.Background(), Input: input, Mock: mock}
	}

	s.Run("publishes a new error", func() {
		sut := makeSut()

		sut.Mock.Publisher.
			On("Handle", *sut.Input).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})
}
