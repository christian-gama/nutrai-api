package consumer_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/notify/app/consumer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
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
		Consumer *messageMock.Consumer[command.SaveUserInput]
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut  consumer.SendWelcomeHandler
		Mock *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer[command.SaveUserInput](s.T()),
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
		Consumer *messageMock.Consumer[command.SaveUserInput]
		Mailer   *mailerMock.Mailer
	}

	type Sut struct {
		Sut   consumer.SendWelcomeHandler
		Mock  *Mock
		Input *command.SaveUserInput
		Ctx   context.Context
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Consumer: messageMock.NewConsumer[command.SaveUserInput](s.T()),
			Mailer:   mailerMock.NewMailer(s.T()),
		}

		sut := consumer.NewSendWelcomeHandler(mock.Consumer, mock.Mailer)

		return &Sut{
			Sut:   sut,
			Mock:  mock,
			Ctx:   context.Background(),
			Input: fake.SaveUserInput(),
		}
	}

	s.Run("should save a new mailer", func() {
		sut := makeSut()

		sut.Mock.Mailer.
			On("Send", sut.Ctx, mock.Anything).
			Return(nil)

		sut.Sut.ConsumerHandler(*sut.Input)

		sut.Mock.Mailer.AssertCalled(
			s.T(),
			"Send",
			sut.Ctx,
			mock.MatchedBy(func(input *mail.Mail) bool {
				return input.To[0].Email == sut.Input.Email.String()
			}),
		)
	})
}
