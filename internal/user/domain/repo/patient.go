package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
)

// SavePatientInput is the input for the Save method.
type SavePatientInput struct {
	Patient *patient.Patient
}

// AllPatientsInput is the input for the All method.
type AllPatientsInput struct {
	querying.Filterer
	querying.Sorter
	querying.Paginator
	querying.Preloader
}

// FindPatientInput is the input for the Find method.
type FindPatientInput struct {
	ID value.ID
	querying.Filterer
	querying.Preloader
}

// DeletePatientInput is the input for the Delete method.
type DeletePatientInput struct {
	IDs []value.ID
}

// UpdatePatientInput is the input for the Update method.
type UpdatePatientInput struct {
	Patient *patient.Patient
	ID      value.ID
}

// Patient is the interface that wraps the basic Patient methods.
type Patient interface {
	// All returns all patients.
	All(ctx context.Context, input AllPatientsInput) (*querying.PaginationOutput[*patient.Patient], error)

	// Delete deletes the patient with the given id.
	Delete(ctx context.Context, input DeletePatientInput) error

	// Find returns the patient with the given id.
	Find(ctx context.Context, input FindPatientInput) (*patient.Patient, error)

	// Save saves the given patient.
	Save(ctx context.Context, input SavePatientInput) (*patient.Patient, error)

	// Update updates the given patient.
	Update(ctx context.Context, input UpdatePatientInput) error
}
