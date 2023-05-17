package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/user/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	"gorm.io/gorm"
)

type PatientDeps struct {
	Patient *patient.Patient
}

func SavePatient(db *gorm.DB, deps *PatientDeps) *PatientDeps {
	if deps == nil {
		deps = &PatientDeps{}
	}

	patient := deps.Patient
	if patient == nil {
		patient = fake.Patient()

		// Ensure that IDs are not set, so that the database can generate them.
		patient.User.ID = 0
		patient.ID = 0

		patient, err := persistence.NewSQLPatient(db).
			Save(context.Background(), repo.SavePatientInput{
				Patient: patient,
			})
		if err != nil {
			panic(fmt.Errorf("could not create patient: %w", err))
		}

		deps.Patient = patient
	}

	return deps
}
