package middleware

import "github.com/christian-gama/nutrai-api/internal/auth/app/query"

func MakeAuth() Auth {
	return NewAuth(query.MakeAuthHandler())
}
