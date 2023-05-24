package patient_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type AllergyTestSuite struct {
	suite.Suite
}

func TestAllergySuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllergyTestSuite))
}

func (s *AllergyTestSuite) TestNewAllergy() {
	type Sut struct {
		Sut  func() (*patient.Allergy, error)
		Data *patient.Allergy
	}

	makeSut := func() *Sut {
		data := fake.Allergy()

		sut := func() (*patient.Allergy, error) {
			return patient.NewAllergy().
				SetID(data.ID).
				SetName(data.Name).
				SetPatientID(data.PatientID).
				Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewAllergy (Error)", func() {
		s.Run("Name", func() {
			s.Run("Should return an error when greater than maximum", func() {
				sut := makeSut()
				sut.Data.Name = value.Allergy(strings.Repeat("a", 256))

				allergy, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(allergy)
			})
		})
	})

	s.Run("TestNewAllergy (Success)", func() {
		s.Run("Should return a allergy when all fields are valid", func() {
			sut := makeSut()

			allergy, err := sut.Sut()

			s.NoError(err)
			s.NotNil(allergy)
			s.Equal(sut.Data.Name, allergy.Name)
		})
	})
}
