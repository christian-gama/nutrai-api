package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/hash"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
)

func MakeChangePasswordHandler() ChangePasswordHandler {
	return NewChangePasswordHandler(
		persistence.MakeSQLUser(),
		hash.MakeHasher(),
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

func MakeSaveUserHandler() SaveUserHandler {
	return NewSaveUserHandler(
		persistence.MakeSQLUser(),
		hash.MakeHasher(),
	)
}
