package query

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// FindUserInput is the input data of the user query.
type FindUserOutput struct {
	ID    sharedvalue.ID `json:"id"`
	Email value.Email    `json:"email"`
	Name  value.Name     `json:"name"`
}
