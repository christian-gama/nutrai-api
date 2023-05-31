package consumer_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/notify/app/consumer"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
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
		Consumer *messageMock.Consumer[user.User]
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut  consumer.SendWelcomeHandler
		Mock *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer[user.User](s.T()),
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
		Consumer *messageMock.Consumer[user.User]
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut  consumer.SendWelcomeHandler
		Mock *Mock
		Ctx  context.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer[user.User](s.T()),
			Mailer:   mailerMock.NewMailer(s.T()),
		}

		sut := consumer.NewSendWelcomeHandler(mock.Consumer, mock.Mailer)

		return &Sut{Sut: sut, Mock: mock, Ctx: context.Background()}
	}

	s.Run("should save a new mailer", func() {
		sut := makeSut()

		sut.Mock.Mailer.
			On("Send", sut.Ctx, mock.Anything).
			Return(nil)

		sut.Sut.ConsumerHandler(*fake.User())

		sut.Mock.Mailer.AssertExpectations(s.T())
	})
}
