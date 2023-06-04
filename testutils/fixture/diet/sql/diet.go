package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/patient/sql"
	"gorm.io/gorm"
)

type DietDeps struct {
	Diet    *diet.Diet
	Patient *patient.Patient
}

func SaveDiet(db *gorm.DB, deps *DietDeps) *DietDeps {
	if deps == nil {
		deps = &DietDeps{}
	}

	if deps.Patient == nil {
		deps.Patient = fixture.SavePatient(db, nil).Patient
	}

	diet := deps.Diet
	if diet == nil {
		diet = fake.Diet()
	}

	diet.ID = deps.Patient.ID
	diet.PatientID = deps.Patient.ID
	diet, err := persistence.NewSQLDiet(db).
		Save(context.Background(), repo.SaveDietInput{
			Diet: diet,
		})
	if err != nil {
		panic(fmt.Errorf("could not create diet: %w", err))
	}

	deps.Diet = diet

	return deps
}
