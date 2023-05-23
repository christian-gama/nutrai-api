package jwt

import "github.com/christian-gama/nutrai-api/pkg/errutil"

var (
	// ErrMissingAuthorizationHeader is the error returned when the authorization header is missing.
	ErrMissingAuthorizationHeader = errutil.NewErrRequired("header:authorization")

	// ErrInvalidAuthorizationHeader is the error returned when the authorization header is invalid.
	ErrInvalidAuthorizationHeader = errutil.NewErrInvalid(
		"header:authorization",
		"it's expected to have a valid bearer token as 'Bearer <token>'",
	)
)
