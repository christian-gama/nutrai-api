package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// FindUserInput is the input data of the user query.
type FindUserOutput struct {
	ID    coreValue.ID `json:"id" faker:"uint"`
	Email value.Email  `json:"email" faker:"email"`
	Name  value.Name   `json:"name" faker:"name"`
}
