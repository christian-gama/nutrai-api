package consumer_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/notify/app/consumer"
	messageMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/message"
	mailerMock "github.com/christian-gama/nutrai-api/testutils/mocks/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type SendWelcomeSuite struct {
	suite.Suite
}

func TestSendWelcomeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SendWelcomeSuite))
}

func (s *SendWelcomeSuite) TestHandle() {
	type Mock struct {
		Consumer *messageMock.Consumer
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut  consumer.SendWelcomeHandler
		Mock *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer(s.T()),
			Mailer:   mailerMock.NewMailer(s.T()),
		}

		sut := consumer.NewSendWelcomeHandler(mock.Consumer, mock.Mailer)

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

func (s *SendWelcomeSuite) TestConsumerHandler() {
	type Mock struct {
		Consumer *messageMock.Consumer
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut  consumer.SendWelcomeHandler
		Mock *Mock
		Ctx  context.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer(s.T()),
			Mailer:   mailerMock.NewMailer(s.T()),
		}

		sut := consumer.NewSendWelcomeHandler(mock.Consumer, mock.Mailer)

		return &Sut{Sut: sut, Mock: mock, Ctx: context.Background()}
	}

	s.Run("should save a new mailer", func() {
		sut := makeSut()

		body, _ := json.Marshal(&user.User{})

		sut.Mock.Mailer.
			On("Send", sut.Ctx, mock.Anything).
			Return(nil)

		sut.Sut.ConsumerHandler(body)

		sut.Mock.Mailer.AssertExpectations(s.T())
	})

	s.Run("should return an error when body is invalid", func() {
		sut := makeSut()

		body := []byte("invalid body")

		err := sut.Sut.ConsumerHandler(body)

		s.NotNil(err)
	})
}
