package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql"
)

func MakeAuthHandler() AuthHandler {
	return NewAuthHandler(jwt.MakeVerifier(), persistence.MakeSQLUser())
}
