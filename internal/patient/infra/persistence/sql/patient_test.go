package persistence_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql/schema"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	userFixture "github.com/christian-gama/nutrai-api/testutils/fixture/auth/sql"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/patient/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type PatientSuite struct {
	suite.SuiteWithSQLConn
	Patient func(db *gorm.DB) repo.Patient
}

func TestPatientSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(PatientSuite))
}

func (s *PatientSuite) SetupTest() {
	s.Patient = func(db *gorm.DB) repo.Patient {
		return persistence.NewSQLPatient(db)
	}
}

func (s *PatientSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SavePatientInput) (*patient.Patient, error)
		Ctx   context.Context
		Input repo.SavePatientInput
	}

	makeSut := func(db *gorm.DB) Sut {
		patient := fake.Patient()

		input := repo.SavePatientInput{
			Patient: patient,
		}

		sut := s.Patient(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
		}
	}

	s.Run("Should create a new patient", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := userFixture.SaveUser(db, nil)
		sut.Input.Patient.ID = userDeps.User.ID
		patient, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(patient.ID, "Should have an ID")
	})

	s.Run("Should return an error when the patient already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		userDeps := userFixture.SaveUser(db, nil)
		sut.Input.Patient.ID = userDeps.User.ID
		_, err := sut.Sut(sut.Ctx, sut.Input)
		s.NoError(err)
		s.SQLRecordExist(db, &schema.Patient{})

		_, err = sut.Sut(sut.Ctx, sut.Input)
		s.Error(err)
	})

	s.Run("Should return an error when the user does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.Patient.ID = 404_404_404
		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}

func (s *PatientSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeletePatientInput) error
		Ctx   context.Context
		Input repo.DeletePatientInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.DeletePatientInput{}
		sut := s.Patient(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete a patient", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := fixture.SavePatient(db, nil)

		sut.Input.IDs = []value.ID{patientDeps.Patient.ID}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.SQLRecordDoesNotExist(db, &schema.Patient{})
	})

	s.Run("Should delete nothing if the patient ID is invalid", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := fixture.SavePatient(db, nil)

		sut.Input.IDs = []value.ID{patientDeps.Patient.ID + 1}

		err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.SQLRecordExist(db, &schema.Patient{})
	})
}

func (s *PatientSuite) TestFind() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.FindPatientInput,
		) (*patient.Patient, error)
		Ctx   context.Context
		Input repo.FindPatientInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.FindPatientInput{
			ID: 0,
		}
		sut := s.Patient(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find a patient", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := fixture.SavePatient(db, nil)

		sut.Input.ID = patientDeps.Patient.ID

		patient, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(patient.ID, patientDeps.Patient.ID, "Should have the same ID")
	})

	s.Run("Should return an error if the patient does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}

func (s *PatientSuite) TestAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.AllPatientsInput,
		) (*queryer.PaginationOutput[*patient.Patient], error)
		Ctx   context.Context
		Input repo.AllPatientsInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.AllPatientsInput{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Patient(db).All

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find all patients", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.SavePatient(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "Should have a valid id")
		s.Equal(length, result.Total, "Should return %d total", length)
		s.Len(result.Results, length, "Should return %d results", length)
	})

	s.Run("Should return the correct patients using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := fixture.SavePatient(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.SavePatient(db, nil)
		}

		sut.Input.Filterer = sut.Input.Filterer.Add(
			"age",
			querying.EqOperator,
			patientDeps.Patient.Age,
		)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(result.Results[0].ID, patientDeps.Patient.ID, "Should have the same ID")
		s.Equal(1, result.Total, "Should return only one patient")
		s.Len(result.Results, 1, "Should return only one patient")
	})

	s.Run("Should return the correct patients using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SavePatient(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(
			int(result.Results[1].ID),
			int(result.Results[2].ID),
			"Should have the correct order",
		)
	})

	s.Run("Should return the correct patients using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SavePatient(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Greater(
			int(result.Results[2].ID),
			int(result.Results[1].ID),
			"Should have the correct order",
		)
	})

	s.Run("Should return the correct patients using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		patients := make([]*patient.Patient, 0)
		for i := 0; i < 3; i++ {
			patientDeps := fixture.SavePatient(db, nil)
			patients = append(patients, patientDeps.Patient)
		}

		sut.Input.Paginator = sut.Input.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Equal(3, result.Total, "Should return the correct total")
		s.Len(result.Results, 1, "Should return the correct number of patients")
		s.Equal(int(patients[0].ID), int(result.Results[0].ID), "Should return the correct patient")
	})
}

func (s *PatientSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.UpdatePatientInput,
		) error
		Ctx   context.Context
		Input repo.UpdatePatientInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.UpdatePatientInput{
			Patient: fake.Patient(),
			ID:      1,
		}
		sut := s.Patient(db).Update

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should update a patient", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := fixture.SavePatient(db, nil)

		*sut.Input.Patient = *patientDeps.Patient
		sut.Input.Patient.Age = 50
		sut.Input.ID = patientDeps.Patient.ID

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Require().NoError(err)
		s.HasChanged(patientDeps.Patient, sut.Input.Patient)
	})
}
