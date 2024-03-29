package patient_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
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
			return data.Validate()
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
		})

		s.Run("HeightM", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.HeightM = 0

				patient, err := sut.Sut()

				s.ErrorAsRequired(err)
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

	s.Run("Allergies", func() {
		s.Run("Should return an error when name is empty", func() {
			sut := makeSut()

			sut.Data.Allergies = []*patient.Allergy{fake.Allergy().SetName("")}

			patient, err := sut.Sut()

			s.ErrorAsRequired(err)
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
			s.Equal(sut.Data.Allergies, patient.Allergies, "should have the same allergies")
			s.Equal(sut.Data.Allergies[0].PatientID, patient.ID, "should have the same patient id")
		})

		s.Run("Should append new allergies and remove the old ones", func() {
			sut := makeSut()

			p, err := sut.Sut()
			s.Require().NoError(err)

			allergies := []*patient.Allergy{fake.Allergy().SetName("new allergy")}
			p = p.SetAllergies(allergies...)

			s.Equal(1, len(p.Allergies), "should have 1 allergy")
			s.Equal(p.ID, p.Allergies[0].PatientID, "should have the same patient id")
			s.Equal(allergies[0].Name, p.Allergies[0].Name, "should have the same allergy name")
		})

		s.Run("Should append new allergies and keep the old ones", func() {
			sut := makeSut()

			p, err := sut.Sut()
			s.Require().NoError(err)
			originalAllergiesLen := len(p.Allergies)

			allergy := []*patient.Allergy{fake.Allergy().SetName("new allergy")}
			p = p.SetAllergies(append(p.Allergies, allergy...)...)

			s.Equal(originalAllergiesLen+1, len(p.Allergies), "should have 1 allergy")
		})
	})
}
