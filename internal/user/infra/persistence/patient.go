package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/manager"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence/schema"
	"gorm.io/gorm"
)

// NewPatient returns a new Patient.
func NewPatient(db *gorm.DB) repo.Patient {
	return &patientImpl{
		manager: manager.NewManager[patient.Patient, schema.Patient](db),
	}
}

// patientImpl is the implementation of repo.Patient.
type patientImpl struct {
	manager *manager.Manager[patient.Patient, schema.Patient]
}

// All implements repo.Patient.
func (p *patientImpl) All(ctx context.Context, input repo.AllPatientsInput) (*querying.PaginationOutput[*patient.Patient], error) {
	return p.manager.All(ctx, manager.AllInput[patient.Patient]{Filterer: input.Filterer, Paginator: input.Paginator, Sorter: input.Sorter, Preloader: input.Preloader})
}

// Delete implements repo.Patient.
func (p *patientImpl) Delete(ctx context.Context, input repo.DeletePatientInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[patient.Patient]{IDs: input.IDs})
}

// Find implements repo.Patient.
func (p *patientImpl) Find(ctx context.Context, input repo.FindPatientInput) (*patient.Patient, error) {
	return p.manager.Find(ctx, manager.FindInput[patient.Patient]{ID: input.ID, Preloader: input.Preloader})
}

// Save implements repo.Patient.
func (p *patientImpl) Save(ctx context.Context, input repo.SavePatientInput) (*patient.Patient, error) {
	return p.manager.Save(ctx, manager.SaveInput[patient.Patient]{Model: input.Patient})
}

// Update implements repo.Patient.
func (p *patientImpl) Update(ctx context.Context, input repo.UpdatePatientInput) error {
	return p.manager.Update(ctx, manager.UpdateInput[patient.Patient]{Model: input.Patient, ID: input.ID})
}
