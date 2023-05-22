package jwt

import "github.com/christian-gama/nutrai-api/pkg/errutil"

var (
	// ErrMissingAuthorizationHeader is the error returned when the authorization header is missing.
	ErrMissingAuthorizationHeader = errutil.NewErrRequired("authorization header")

	// ErrInvalidAuthorizationHeader is the error returned when the authorization header is invalid.
	ErrInvalidAuthorizationHeader = errutil.NewErrInvalid(
		"authorization header",
		"it's not a bearer token",
	)
)
