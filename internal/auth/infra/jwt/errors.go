package jwt

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

var (
	// ErrMissingAuthorizationHeader is the error returned when the authorization header is missing.
	ErrMissingAuthorizationHeader = errors.Unauthorized("missing header: Authorization")

	// ErrInvalidAuthorizationHeader is the error returned when the authorization header is invalid.
	ErrInvalidAuthorizationHeader = errors.Unauthorized(
		"Authorization token expects a valid bearer token",
	)

	ErrInvalidToken = errors.Unauthorized("invalid token")
)
