package consumer_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/internal/exception/app/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/domain/model/exception"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type RecoverySuite struct {
	suite.Suite
}

func TestRecoverySuite(t *testing.T) {
	suite.RunUnitTest(t, new(RecoverySuite))
}

func (s *RecoverySuite) TestHandle() {
	type Mock struct {
		Consumer      *messageMock.Consumer[command.RecoveryInput]
		ExceptionRepo *repoMock.Exception
	}

	type Sut struct {
		Sut  consumer.RecoveryHandler
		Mock *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer:      messageMock.NewConsumer[command.RecoveryInput](s.T()),
			ExceptionRepo: repoMock.NewException(s.T()),
		}

		sut := consumer.NewRecoveryHandler(mock.Consumer, mock.ExceptionRepo)

		return &Sut{Sut: sut, Mock: mock}
	}

	s.Run("consumes a new error", func() {
		sut := makeSut()

		sut.Mock.Consumer.
			On("Handle", mock.Anything).
			Return(nil)

		sut.Sut.Handle()

		sut.Mock.Consumer.AssertExpectations(s.T())
	})
}

func (s *RecoverySuite) TestConsumerHandler() {
	type Mock struct {
		Consumer      *messageMock.Consumer[command.RecoveryInput]
		ExceptionRepo *repoMock.Exception
	}

	type Sut struct {
		Sut   consumer.RecoveryHandler
		Mock  *Mock
		Input command.RecoveryInput
		Ctx   context.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer:      messageMock.NewConsumer[command.RecoveryInput](s.T()),
			ExceptionRepo: repoMock.NewException(s.T()),
		}

		sut := consumer.NewRecoveryHandler(mock.Consumer, mock.ExceptionRepo)

		return &Sut{
			Sut:  sut,
			Mock: mock,
			Ctx:  context.Background(),
			Input: command.RecoveryInput{
				Message: "test",
				Stack:   "test",
			},
		}
	}

	s.Run("should save a new exception", func() {
		sut := makeSut()

		sut.Mock.ExceptionRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fake.Exception(), nil)

		sut.Sut.ConsumerHandler(sut.Input)

		sut.Mock.ExceptionRepo.AssertCalled(
			s.T(),
			"Save",
			sut.Ctx,
			mock.MatchedBy(func(input repo.SaveExceptionInput) bool {
				return input.Exception.Message == sut.Input.Message &&
					input.Exception.Stack == sut.Input.Stack
			}),
		)
	})
}
