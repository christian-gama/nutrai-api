package query

import (
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// FindByEmailInput is the query to find a user by email.
type FindByEmailInput struct {
	Email userValue.Email `json:"email" validate:"required,email" faker:"email"`
}

// FindByEmailOutput is the output of the query to find a user by email.
type FindByEmailOutput struct {
	ID    coreValue.ID    `json:"id"`
	Email userValue.Email `json:"email"`
	Name  userValue.Name  `json:"name"`
}
