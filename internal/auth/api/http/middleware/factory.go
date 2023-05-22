package middleware

import "github.com/christian-gama/nutrai-api/internal/auth/app/query"

func MakeAuthHandler() AuthHandler {
	return NewAuthHandler(query.MakeAuthHandler())
}
