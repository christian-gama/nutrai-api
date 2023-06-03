package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/event"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/hash"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	redisPersistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
	sqlPersistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
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
		publisher.MakePublisher[user.User](
			publisher.WithExchangeName(event.User),
			publisher.WithRoutingKey(event.SaveUser),
		),
	)
}

func MakeLogoutHandler() LogoutHandler {
	return NewLogoutHandler(redisPersistence.MakeRedisToken(), jwt.MakeRefreshVerifier())
}

func MakeLogoutAllHandler() LogoutAllHandler {
	return NewLogoutAllHandler(redisPersistence.MakeRedisToken())
}
