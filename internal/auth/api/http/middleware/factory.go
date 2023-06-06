// Internal middlewares for the routes. It's important to note that the middlewares here are
// different from the middlewares from the api/http/middlewares, because the middlewares here
// are used by the routes internally.

package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
)

func MakeJwtAuth() JwtAuth {
	return NewJwtAuth(query.MakeJwtAuthHandler())
}

func MakeApiKey() ApiKeyAuth {
	return NewApiKeyAuth(query.MakeApiKeyAuthHandler())
}
