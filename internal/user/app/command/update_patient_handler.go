package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// UpdatePatientHandler represents the UpdatePatient command.
type UpdatePatientHandler = command.Handler[*UpdatePatientInput]

// updatePatientHandlerImpl represents the implementation of the UpdatePatient command.
type updatePatientHandlerImpl struct {
	repo.Patient
}

// NewUpdatePatientHandler returns a new Update instance.
func NewUpdatePatientHandler(p repo.Patient) UpdatePatientHandler {
	return &updatePatientHandlerImpl{p}
}

// Handle implements command.Handler.
func (c *updatePatientHandlerImpl) Handle(ctx context.Context, input *UpdatePatientInput) error {
	savedPatient, err := c.Find(ctx, repo.FindPatientInput{ID: value.ID(input.ID)}, "User")
	if err != nil {
		return err
	}

	patient, err := convert.ToValidModel(savedPatient, input)
	if err != nil {
		return err
	}

	return c.Update(ctx, repo.UpdatePatientInput{Patient: patient, ID: value.ID(input.ID)})
}
