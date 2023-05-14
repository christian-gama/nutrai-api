package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
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
func NewSavePatientHandler(r repo.Patient, h service.HashPasswordHandler) SavePatientHandler {
	if r == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	if h == nil {
		panic(errors.New("service.HashPasswordHandler cannot be nil"))
	}

	return &savePatientHandlerImpl{r, h}
}

// Handle implements command.Handler.
func (c *savePatientHandlerImpl) Handle(ctx context.Context, input *SavePatientInput) error {
	hashPasswordOutput, err := c.HashPasswordHandler.Handle(
		ctx, &service.HashPasswordInput{Password: input.User.Password},
	)
	if err != nil {
		return err
	}

	user, err := user.NewBuilder().
		SetName(input.User.Name).
		SetEmail(input.User.Email).
		SetPassword(hashPasswordOutput.Password).
		Build()
	if err != nil {
		return err
	}

	patient, err := patient.NewBuilder().
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
		SetUser(user).
		Build()
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SavePatientInput{Patient: patient}))
}
