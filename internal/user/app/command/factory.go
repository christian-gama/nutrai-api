package command

import (
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/infra/hash"
	persistence "github.com/christian-gama/nutrai-api/internal/user/infra/persistence/sql"
)

func MakeChangePasswordHandler() ChangePasswordHandler {
	return NewChangePasswordHandler(
		persistence.MakeSQLUser(),
		service.MakeHashPasswordHandler(),
	)
}

func MakeCheckCredentialsHandler() CheckCredentialsHandler {
	return NewCheckCredentialsHandler(
		persistence.MakeSQLUser(),
		hash.MakeHasher(),
	)
}

func MakeDeleteUserHandler() DeleteUserHandler {
	return NewDeleteUserHandler(
		persistence.MakeSQLUser(),
	)
}

func MakeSavePatientHandler() SavePatientHandler {
	return NewSavePatientHandler(
		persistence.MakeSQLPatient(),
		service.MakeHashPasswordHandler(),
	)
}

func MakeUpdatePatientHandler() UpdatePatientHandler {
	return NewUpdatePatientHandler(
		persistence.MakeSQLPatient(),
	)
}
