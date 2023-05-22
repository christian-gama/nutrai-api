package store

import "github.com/christian-gama/nutrai-api/pkg/errutil"

var (
	ErrUserNotFound  = errutil.NewErrNotFound("user")
	ErrUserIsInvalid = errutil.NewErrInternal("user is invalid")
)
