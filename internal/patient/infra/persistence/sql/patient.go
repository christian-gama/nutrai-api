package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/sqlerr"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql/schema"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"gorm.io/gorm"
)

// patientSQLImpl is the SQL implementation of repo.Patient.
type patientSQLImpl struct {
	manager *manager.Manager[patient.Patient, schema.Patient]
}

// NewSQLPatient returns a new Patient.
func NewSQLPatient(db *gorm.DB) repo.Patient {
	errutil.MustBeNotEmpty("gorm.DB", db)

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
	db := p.manager.DB.WithContext(ctx)
	schema := convert.FromModel(&schema.Patient{}, &input.Patient)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.
			Unscoped().
			Model(&schema).
			Association("Allergies").
			Unscoped().
			Replace(input.Patient.Allergies); err != nil {
			return sqlerr.Error(err, input.Patient)
		}

		if err := db.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Model(&schema).
			Where("id = ?", input.ID).
			Updates(&schema).
			Error; err != nil {
			return sqlerr.Error(err, input.Patient)
		}

		return nil
	}); err != nil {
		return sqlerr.Error(err, input.Patient)
	}

	return nil
}
