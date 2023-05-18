package jwt

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// Subject is the data that will be encoded in the JWT token.
type Subject struct {
	ID    coreValue.ID `json:"id" faker:"uint"`
	Email value.Email  `json:"email" faker:"email"`
}

// Generator is the interface that wraps the Generate method.
type Generator interface {
	// Generate generates a new JWT token.
	Generate(data *Subject) (value.Token, error)
}
