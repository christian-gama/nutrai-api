// Internal middlewares for the routes. It's important to note that the middlewares here are
// different from the middlewares from the api/http/middlewares, because the middlewares here
// are used by the routes internally.

package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
)

func MakeAuthJwt() AuthJwt {
	return NewAuthJwt(query.MakeAuthJwtHandler())
}

func MakeAuthApiKey() AuthApiKey {
	return NewAuthApiKey(query.MakeAuthApiKeyHandler())
}
