package gpt_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	value "github.com/christian-gama/nutrai-api/internal/gpt/domain/value/gpt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/gpt/domain/model/gpt"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ModelTestSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ModelTestSuite))
}

func (s *ModelTestSuite) TestNewModel() {
	type Sut struct {
		Sut  func() (*gpt.Model, error)
		Data *gpt.Model
	}

	makeSut := func() *Sut {
		data := fake.Model()

		sut := func() (*gpt.Model, error) {
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewModel (Error)", func() {
		s.Run("Name", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Name = ""

				model, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(model)
			})
		})

		s.Run("MaxTokens", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.MaxTokens = 0

				model, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(model)
			})
		})

		s.Run("Temperature", func() {
			s.Run("Should return an error when is less than 0", func() {
				sut := makeSut()
				sut.Data.Temperature = value.Temperature(-1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})

			s.Run("Should return an error when is greater than 1", func() {
				sut := makeSut()
				sut.Data.Temperature = value.Temperature(1.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})
		})

		s.Run("TopP", func() {
			s.Run("Should return an error when is less than 0", func() {
				sut := makeSut()
				sut.Data.TopP = value.TopP(-1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})

			s.Run("Should return an error when is greater than 1", func() {
				sut := makeSut()
				sut.Data.TopP = value.TopP(1.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})
		})

		s.Run("N", func() {
			s.Run("Should return an error when is less than 0", func() {
				sut := makeSut()
				sut.Data.N = value.N(-1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})
		})

		s.Run("FrequencyPenalty", func() {
			s.Run("Should return an error when is less than -2", func() {
				sut := makeSut()
				sut.Data.FrequencyPenalty = value.FrequencyPenalty(-2.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})

			s.Run("Should return an error when is greater than 2", func() {
				sut := makeSut()
				sut.Data.FrequencyPenalty = value.FrequencyPenalty(2.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})
		})

		s.Run("PresencePenalty", func() {
			s.Run("Should return an error when is less than -2", func() {
				sut := makeSut()
				sut.Data.PresencePenalty = value.PresencePenalty(-2.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})

			s.Run("Should return an error when is greater than 2", func() {
				sut := makeSut()
				sut.Data.PresencePenalty = value.PresencePenalty(2.1)

				model, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(model)
			})
		})
	})
}
