package query

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
)

func MakeFindByEmailHandler() FindByEmailHandler {
	return NewFindByEmailHandler(persistence.MakeSQLUser())
}

func MakeAuthJwtHandler() AuthJwtHandler {
	return NewAuthJwtHandler(persistence.MakeSQLUser(), jwt.MakeAccessVerifier())
}

func MakeAuthApiKeyHandler() AuthApiKeyHandler {
	return NewAuthApiKeyHandler()
}
