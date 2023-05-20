package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/auth/sql"
	"gorm.io/gorm"
)

type PatientDeps struct {
	Patient *patient.Patient
	User    *user.User
}

func SavePatient(db *gorm.DB, deps *PatientDeps) *PatientDeps {
	if deps == nil {
		deps = &PatientDeps{}
	}

	if deps.User == nil {
		deps.User = fixture.SaveUser(db, nil).User
	}

	patient := deps.Patient
	if patient == nil {
		patient = fake.Patient()
		patient.ID = deps.User.ID

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
