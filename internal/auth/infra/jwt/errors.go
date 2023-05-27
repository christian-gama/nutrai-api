package jwt

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

var (
	// ErrMissingAuthorizationHeader is the error returned when the authorization header is missing.
	ErrMissingAuthorizationHeader = errors.Required("header:authorization")

	// ErrInvalidAuthorizationHeader is the error returned when the authorization header is invalid.
	ErrInvalidAuthorizationHeader = errors.Invalid(
		"header:authorization",
		"it's expected to have a valid bearer token as 'Bearer <token>'",
	)

	ErrInvalidToken = errors.Invalid(
		"token",
		"the token is invalid",
	)
)
