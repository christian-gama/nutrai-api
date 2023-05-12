package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/querying"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
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
	savedPatient, err := c.Find(ctx,
		repo.FindPatientInput{
			ID:        input.ID,
			Preloader: querying.AddPreload("user"),
		},
	)
	if err != nil {
		return err
	}

	user, err := user.NewBuilder().
		SetName(input.User.Name).
		SetEmail(input.User.Email).
		// The password must keep the same. To change the password, refer to the ChangePassword command.
		SetPassword(savedPatient.User.Password).
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

	return c.Update(ctx, repo.UpdatePatientInput{Patient: patient, ID: input.ID})
}
