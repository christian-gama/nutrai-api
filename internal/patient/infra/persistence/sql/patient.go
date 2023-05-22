package persistence

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql/schema"
	"gorm.io/gorm"
)

// patientSQLImpl is the SQL implementation of repo.Patient.
type patientSQLImpl struct {
	manager *manager.Manager[patient.Patient, schema.Patient]
}

// NewSQLPatient returns a new Patient.
func NewSQLPatient(db *gorm.DB) repo.Patient {
	if db == nil {
		panic(errors.New("db cannot be nil"))
	}

	return &patientSQLImpl{
		manager: manager.NewManager[patient.Patient, schema.Patient](db),
	}
}

// All implements repo.Patient.
func (p *patientSQLImpl) All(
	ctx context.Context,
	input repo.AllPatientsInput,
) (*queryer.PaginationOutput[*patient.Patient], error) {
	return p.manager.All(ctx,
		manager.AllInput[patient.Patient]{
			Filterer:  input.Filterer,
			Paginator: input.Paginator,
			Sorter:    input.Sorter,
			Preloader: input.Preloader,
		},
	)
}

// Delete implements repo.Patient.
func (p *patientSQLImpl) Delete(ctx context.Context, input repo.DeletePatientInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[patient.Patient]{IDs: input.IDs})
}

// Find implements repo.Patient.
func (p *patientSQLImpl) Find(
	ctx context.Context,
	input repo.FindPatientInput,
) (*patient.Patient, error) {
	return p.manager.Find(ctx,
		manager.FindInput[patient.Patient]{
			ID:        input.ID,
			Preloader: input.Preloader,
		},
	)
}

// Save implements repo.Patient.
func (p *patientSQLImpl) Save(
	ctx context.Context,
	input repo.SavePatientInput,
) (*patient.Patient, error) {
	return p.manager.Save(ctx, manager.SaveInput[patient.Patient]{Model: input.Patient})
}

// Update implements repo.Patient.
func (p *patientSQLImpl) Update(ctx context.Context, input repo.UpdatePatientInput) error {
	return p.manager.Update(ctx,
		manager.UpdateInput[patient.Patient]{
			Model: input.Patient,
			ID:    input.ID,
		},
	)
}