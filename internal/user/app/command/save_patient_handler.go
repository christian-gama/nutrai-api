package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// SavePatientHandler represents the SavePatient command.
type SavePatientHandler = command.Handler[*SavePatientInput]

// savePatientHandlerImpl represents the implementation of the SavePatient command.
type savePatientHandlerImpl struct {
	repo.Patient
	service.HashPasswordHandler
}

// NewSavePatientHandler returns a new Save instance.
func NewSavePatientHandler(p repo.Patient, h service.HashPasswordHandler) SavePatientHandler {
	return &savePatientHandlerImpl{p, h}
}

// Handle implements command.Handler.
func (c *savePatientHandlerImpl) Handle(ctx context.Context, input *SavePatientInput) error {
	hashPasswordOutput, err := c.HashPasswordHandler.Handle(ctx, &service.HashPasswordInput{Password: input.User.Password})
	if err != nil {
		return err
	}
	input.User.Password = hashPasswordOutput.HashedPassword

	patient, err := convert.ToValidModel(&patient.Patient{}, input)
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SavePatientInput{Patient: patient}))
}
