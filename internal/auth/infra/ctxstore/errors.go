package ctxstore

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

var (
	ErrUserNotFound = errors.InternalServerError(
		"there is no user in the context - did you forget to set the route to private (Controller.Public = false)?",
	)
	ErrUserIsInvalid = errors.InternalServerError(
		"the user in the context was found, but it's invalid - are you sure you're using the right user.User from auth module?",
	)
)
