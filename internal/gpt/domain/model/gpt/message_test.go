package gpt_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/gpt/domain/model/gpt"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type MessageTestSuite struct {
	suite.Suite
}

func TestMessageSuite(t *testing.T) {
	suite.RunUnitTest(t, new(MessageTestSuite))
}

func (s *MessageTestSuite) TestNewMessage() {
	type Sut struct {
		Sut  func() (*gpt.Message, error)
		Data *gpt.Message
	}

	makeSut := func() *Sut {
		data := fake.Message()

		sut := func() (*gpt.Message, error) {
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewMessage (Error)", func() {
		s.Run("Role", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Role = ""

				message, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(message)
			})

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.Role = "invalid"

				message, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(message)
			})
		})

		s.Run("Content", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Content = ""

				message, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(message)
			})
		})

		s.Run("Model", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Model = nil

				message, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(message)
			})
		})
	})
}
