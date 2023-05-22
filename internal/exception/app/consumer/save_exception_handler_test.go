package consumer_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/exception/app/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/exception/domain/model/exception"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type SaveExceptionSuite struct {
	suite.Suite
}

func TestSaveExceptionSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SaveExceptionSuite))
}

func (s *SaveExceptionSuite) TestHandle() {
	type Mock struct {
		Consumer      *messageMock.Consumer
		ExceptionRepo *repoMock.Exception
	}

	type Sut struct {
		Sut  consumer.SaveExceptionHandler
		Mock *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer:      messageMock.NewConsumer(s.T()),
			ExceptionRepo: repoMock.NewException(s.T()),
		}

		sut := consumer.NewSaveExceptionHandler(mock.Consumer, mock.ExceptionRepo)

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

func (s *SaveExceptionSuite) TestConsumerHandler() {
	type Mock struct {
		Consumer      *messageMock.Consumer
		ExceptionRepo *repoMock.Exception
	}

	type Sut struct {
		Sut  consumer.SaveExceptionHandler
		Mock *Mock
		Ctx  context.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer:      messageMock.NewConsumer(s.T()),
			ExceptionRepo: repoMock.NewException(s.T()),
		}

		sut := consumer.NewSaveExceptionHandler(mock.Consumer, mock.ExceptionRepo)

		return &Sut{Sut: sut, Mock: mock, Ctx: context.Background()}
	}

	s.Run("should save a new exception", func() {
		sut := makeSut()

		body, _ := json.Marshal(&exception.Exception{})

		sut.Mock.ExceptionRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fake.Exception(), nil)

		sut.Sut.ConsumerHandler(body)

		sut.Mock.ExceptionRepo.AssertExpectations(s.T())
	})

	s.Run("should return an error when body is invalid", func() {
		sut := makeSut()

		body := []byte("invalid body")

		err := sut.Sut.ConsumerHandler(body)

		s.NotNil(err)
	})
}