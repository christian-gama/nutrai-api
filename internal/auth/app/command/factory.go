package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/hash"
	redisPersistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
	sqlPersistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
)

func MakeChangePasswordHandler() ChangePasswordHandler {
	return NewChangePasswordHandler(
		sqlPersistence.MakeSQLUser(),
		hash.MakeHasher(),
	)
}

func MakeCheckCredentialsHandler() CheckCredentialsHandler {
	return NewCheckCredentialsHandler(
		sqlPersistence.MakeSQLUser(),
		hash.MakeHasher(),
	)
}

func MakeDeleteMeHandler() DeleteMeHandler {
	return NewDeleteMeHandler(
		sqlPersistence.MakeSQLUser(),
	)
}

func MakeSaveUserHandler() SaveUserHandler {
	return NewSaveUserHandler(
		sqlPersistence.MakeSQLUser(),
		hash.MakeHasher(),
	)
}

func MakeLogoutHandler() LogoutHandler {
	return NewLogoutHandler(redisPersistence.MakeRedisToken())
}
