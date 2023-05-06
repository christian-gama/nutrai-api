package user

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// UserInput is the input to create a new User.
type UserInput struct {
	ID       sharedvalue.ID
	Email    value.Email
	Password value.Password
	Name     value.Name
}
