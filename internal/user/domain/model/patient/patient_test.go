package patient_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type PatientTestSuite struct {
	suite.Suite
}

func TestPatientSuite(t *testing.T) {
	suite.RunUnitTest(t, new(PatientTestSuite))
}

func (s *PatientTestSuite) TestNewPatient() {
	type Sut struct {
		Sut  func() (*patient.Patient, error)
		Data *patient.Patient
	}

	makeSut := func() *Sut {
		data := fake.Patient()

		sut := func() (*patient.Patient, error) {
			return patient.NewBuilder().
				SetAge(data.Age).
				SetHeightM(data.HeightM).
				SetWeightKG(data.WeightKG).
				SetUser(data.User).
				Build()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewPatient (Error)", func() {
		s.Run("WeightKG", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.WeightKG = 0

				patient, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when less than minimum", func() {
				sut := makeSut()
				sut.Data.WeightKG = -1

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when greater than maximum", func() {
				sut := makeSut()
				sut.Data.WeightKG = 1000

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})
		})

		s.Run("HeightM", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.HeightM = 0

				patient, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when less than minimum", func() {
				sut := makeSut()
				sut.Data.HeightM = -1

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when greater than maximum", func() {
				sut := makeSut()
				sut.Data.HeightM = 1000

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})
		})

		s.Run("Age", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Age = 0

				patient, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when less than minimum", func() {
				sut := makeSut()
				sut.Data.Age = -1

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})

			s.Run("Should return an error when greater than maximum", func() {
				sut := makeSut()
				sut.Data.Age = 125

				patient, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(patient)
			})
		})

		s.Run("Should return multiple errors when multiple fields are invalid", func() {
			sut := makeSut()
			sut.Data.ID = 0
			sut.Data.WeightKG = 0
			sut.Data.HeightM = 0
			sut.Data.Age = 0

			patient, err := sut.Sut()

			e := err.(*errutil.Error)
			s.Equal(3, e.Len(), "should have 3 errors")
			s.Nil(patient)
		})
	})

	s.Run("TestNewPatient (Success)", func() {
		s.Run("Should return a patient when all fields are valid", func() {
			sut := makeSut()

			patient, err := sut.Sut()

			s.NoError(err)
			s.NotNil(patient)
			s.Equal(sut.Data.WeightKG, patient.WeightKG, "should have the same weight")
			s.Equal(sut.Data.HeightM, patient.HeightM, "should have the same height")
			s.Equal(sut.Data.Age, patient.Age, "should have the same age")
		})
	})
}
