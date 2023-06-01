package patient_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
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
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

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
