package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
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
		user := SaveUser(db, nil)
		deps.User = user.User
	}

	patient := deps.Patient
	if patient == nil {
		patient = fake.Patient()
		patient.User = deps.User

		patient, err := persistence.NewPatient(db).
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
