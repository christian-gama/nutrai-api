package service_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/gpt/app/service"
	serviceFake "github.com/christian-gama/nutrai-api/testutils/fake/gpt/app/service"
	messageFake "github.com/christian-gama/nutrai-api/testutils/fake/gpt/domain/model/gpt"
	generativeMock "github.com/christian-gama/nutrai-api/testutils/mocks/gpt/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ChatCompletionTestSuite struct {
	suite.Suite
}

func TestChatCompletionSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ChatCompletionTestSuite))
}

func (s *ChatCompletionTestSuite) TestNewChatCompletion() {
	type Mock struct {
		Client *generativeMock.Generative
	}

	type Sut struct {
		Sut   service.ChatCompletion
		Mock  *Mock
		Ctx   context.Context
		Input *service.ChatCompletionInput
	}

	makeSut := func() *Sut {
		mock := &Mock{
			Client: generativeMock.NewGenerative(s.T()),
		}

		sut := service.NewChatCompletion(mock.Client)

		input := serviceFake.ChatCompletionInput()

		return &Sut{
			Sut:   sut,
			Mock:  mock,
			Ctx:   context.Background(),
			Input: input,
		}
	}

	s.Run("TestNewChatCompletion (Error)", func() {
		s.Run("Should return an error if the client returns an error", func() {
			sut := makeSut()
			sut.Mock.Client.On("ChatCompletion", mock.Anything, mock.Anything).
				Return(nil, assert.AnError)

			output, err := sut.Sut.Execute(sut.Ctx, sut.Input)

			assert.Nil(s.T(), output)

			assert.Equal(s.T(), assert.AnError, err)
		})
	})

	s.Run("TestNewChatCompletion (Success)", func() {
		s.Run("Should return a valid output", func() {
			sut := makeSut()
			message := messageFake.Message()
			sut.Mock.Client.On("ChatCompletion", mock.Anything, mock.Anything).
				Return(message, nil)

			output, err := sut.Sut.Execute(sut.Ctx, sut.Input)

			s.Nil(err)

			s.Equal(message.Content.String(), output.Message.Content)
			s.Equal(message.Role.String(), output.Message.Role)
		})
	})
}
