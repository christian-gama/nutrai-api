package query

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
)

func MakeFindByEmailHandler() FindByEmailHandler {
	return NewFindByEmailHandler(persistence.MakeSQLUser())
}

func MakeJwtAuthHandler() JwtAuthHandler {
	return NewJwtAuthHandler(persistence.MakeSQLUser(), jwt.MakeAccessVerifier())
}
